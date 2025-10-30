package user

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type TransferDTO struct {
	From   string
	To     string
	Amount uint64
}

func NewTransferDTO(from, to string, amount uint64) *TransferDTO {
	return &TransferDTO{
		From:	from,
		To:		to,
		Amount: amount,
	}
}

type BlockchainGateway interface {
	PublishUserAction(ctx context.Context, payload *UserGoodAction)
	TransferListener(ctx context.Context, handler func(event TransferDTO)) error
	GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error)
}

type BlockchainEventListener interface {
	TransferListener(ctx context.Context) error
}

type UserProjectBonusDTO struct {
	ProjectName string
	Points      int
}

type UserLoggedDaysDTO struct {
	Date string
}

type UserPresenceInEventDTO struct {
	Name        string
	Date        string
	HasPresence bool
}

type UserBasicInfoDTO struct {
	ID42  string `json:"id42"`  // 42 ID
	Login string `json:"login"` // 42 login
}

type Api42Gateway interface {
	GetRetroactiveBonusProject(userID string) ([]UserProjectBonusDTO, error)
	GetRetroactiveLoggedDays(userID string, startAt string) ([]UserLoggedDaysDTO, error)
	GetBasicUserInfo(ctx context.Context, token string) (*UserBasicInfoDTO, error)
	ExchangeCodeForToken(ctx context.Context, code string) (string, error)
}

type ApiGoogleGateway interface {
	GetRetroativePresencesBy42ID(ID42 string) ([]UserPresenceInEventDTO, error)
}

type UserBenefitPurchaseDTO struct {
	ID                   string						`json:"id" bson:"_id"`
	UserID               string						`json:"user_id" bson:"user_id"`
	Username             string						`json:"username" bson:"username"`
	BenefitID            string						`json:"benefit_id" bson:"benefit_id"`
	BenefitName          string						`json:"benefit_name" bson:"benefit_name"`
	EarnedCoinsUsed      uint64						`json:"earned_coins_used" bson:"earned_coins_used"`
	TransferredCoinsUsed uint64						`json:"transferred_coins_used" bson:"transferred_coins_used"`
	PurchaseStatus       UserBenefitPurchaseStatus	`json:"purchase_status" bson:"purchase_status"`
	RequestedAt          string						`json:"requested_at" bson:"requested_at"`
	UpdatedAt            string						`json:"updated_at" bson:"updated_at"`
}

func NewUserBenefitPurchaseDTO(purchase UserBenefitPurchase) *UserBenefitPurchaseDTO {
	return &UserBenefitPurchaseDTO {
		ID:						purchase.ID.Hex(),
		UserID:					purchase.UserID,
		Username:				purchase.Username,
		BenefitID:				purchase.BenefitID.Hex(),
		BenefitName:			purchase.BenefitName,
		EarnedCoinsUsed:		purchase.EarnedCoinsUsed,
		TransferredCoinsUsed:	purchase.TransferredCoinsUsed,
		PurchaseStatus:			purchase.PurchaseStatus,
		RequestedAt:			purchase.RequestedAt.String(),
		UpdatedAt:				purchase.UpdatedAt.String(),
	}
}
