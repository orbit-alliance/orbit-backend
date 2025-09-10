package user

import (
	"os"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
)

func (u *User) Register42(user *User) (*User42Registered, error) {
	return NewUser42Registered(user), nil
}

func (u *User) ChangeWalletAddress(newWallet string, coinsStatus coin.CoinStatus) (*UserWalletChanged, error) {
	if u.Wallet == newWallet {
		return nil, ErrSameWalletAddress
	}

	u.Wallet = newWallet
	u.CoinStatus = coinsStatus

	return NewUserWalletChanged(u), nil
}

// Deducts the transferred amount from the total balance, prioritizing the transferred balance and then the balance earned by shares
func (u *User) SendCoins(to *User, amount uint64) (*SentCoins, error) {
	totalBalance := u.CoinStatus.EarnedByTransfer + u.CoinStatus.EarnedByActions
	if totalBalance < amount {
		return nil, ErrInsufficientBalance
	}
	transferUsed := min(u.CoinStatus.EarnedByTransfer, amount)
	u.CoinStatus.EarnedByTransfer -= transferUsed

	remaining := amount - transferUsed
	u.CoinStatus.EarnedByActions -= remaining

	return NewSentCoins(u, to, amount), nil
}

func (u *User) ReceiveCoins(from *User, amount uint64) (*ReceivedCoins, error) {
	u.CoinStatus.EarnedByActions += amount

	return NewReceivedCoins(from, u, amount), nil
}

func (u *User) doGoodAction(goodAction *UserGoodAction) (*DidGoodAction, error) {

	if goodAction.UserID != u.ID {
		return nil, ErrUserNotAuthorized
	}

	u.CoinStatus.EarnedByActions += uint64(goodAction.RewardAmount)
	return NewDidGoodAction(u, goodAction), nil
}

func (u *User) DoBonusProject(newProject, oldProject *UserProject, goodAction coin.GoodAction) (*DidGoodAction, *UserGoodAction, error) {

	if oldProject == nil {
		userGoodAction := NewUserGoodAction(shared.NewMongoID(), u.ID, u.Wallet, u.Username, goodAction.ID, goodAction.Name, int64(newProject.MaxScore-DefaultMaxScore), time.Now())
		didGoodAction, err := u.doGoodAction(userGoodAction)
		return didGoodAction, userGoodAction, err
	}

	return u.retryProject(newProject, oldProject, goodAction)
}

func (u *User) DoFrequencyReward(newLastLogin string, newStreak int, goodAction coin.GoodAction) (*DidGoodAction, *UserGoodAction, bool, error) {
	// Verificar se a data atual é posterior à FIRE_START_DATE
	fireStartDateStr := os.Getenv("FIRE_START_DATE")
	if fireStartDateStr == "" {
		fireStartDateStr = "2025-06-10T00:00:00Z" // valor padrão
	}

	shouldIncrementRewardedDays := false
	fireStartDate, err := time.Parse(time.RFC3339, fireStartDateStr)
	if err == nil && time.Now().After(fireStartDate) {
		shouldIncrementRewardedDays = true
	}

	u.LastLoginIn42 = newLastLogin
	u.CurrentStreak = newStreak
	userGoodAction := NewUserGoodAction(shared.NewMongoID(), u.ID, u.Wallet, u.Username, goodAction.ID, goodAction.Name, int64(goodAction.RewardAmount), time.Now())
	didGoodAction, err := u.doGoodAction(userGoodAction)

	return didGoodAction, userGoodAction, shouldIncrementRewardedDays, err
}

func (u *User) retryProject(newProject, olderProject *UserProject, goodAction coin.GoodAction) (*DidGoodAction, *UserGoodAction, error) {

	if olderProject.MaxScore < newProject.MaxScore {
		userGoodAction := NewUserGoodAction(shared.NewMongoID(), u.ID, u.Wallet, u.Username, goodAction.ID, goodAction.Name, int64(newProject.MaxScore-olderProject.MaxScore-DefaultMaxScore), time.Now())
		didGoodAction, err := u.doGoodAction(userGoodAction)
		return didGoodAction, userGoodAction, err
	}
	return nil, nil, nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
