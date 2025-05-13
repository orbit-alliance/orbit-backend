package user

import (
	"time"

	gateway_42 "github.com/orbit-alliance/orbit-backend/internal/infra/gateway"
)

func (u *User) EarnCoins(amount uint64) {
	u.CoinStatus.EarnedByActions += amount
}

// Deducts the transferred amount from the total balance, prioritizing the transferred balance and then the balance earned by shares
func (u *User) TransferCoins(amount uint64) (*TransferedCoins, error) {
	totalBalance := u.CoinStatus.EarnedByTransfer + u.CoinStatus.EarnedByActions
	if totalBalance < amount {
		return nil, ErrInsufficientBalance
	}

	u.CoinStatus.Transferred += amount

	transferUsed := min(u.CoinStatus.EarnedByTransfer, amount)
	u.CoinStatus.EarnedByTransfer -= transferUsed

	remaining := amount - transferUsed
	u.CoinStatus.EarnedByActions -= remaining

	return NewTransferedCoins(u, u, amount), nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func GetRetroativeBonusProject(userID string) (map[string]int, error) {
	token, err := gateway_42.GetToken()
	bonusList := make(map[string]int)
	apiData, err := gateway_42.GetUserByID(userID, token)

	if err != nil {
		return nil, ErrFailToGetUserIn42
	}

	for _, project := range apiData.ProjectsUsers {
		if project.FinalMark > 100 {
			bonusList[project.Project.Name] = project.FinalMark
		}
	}

	return bonusList, nil
}

func GetRetroativeLoggedDays(userID string, startAt string) (map[string]string, error) {
	token, err := gateway_42.GetToken()
	endAt := time.Now().UTC().Format(time.RFC3339Nano)

	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	daylist, err := gateway_42.GetLocationByUserID(startAt, endAt, userID, token)

	if err != nil {
		return nil, ErrFailToGetLocationIn42
	}

	return daylist, nil
}
