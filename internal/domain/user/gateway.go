package user

type UserProjectBonusDTO struct {
	ProjectName string
	Points      int
}

type UserLoggedDaysDTO struct {
	Date string
}

type Gateway interface {
	GetRetroactiveBonusProject(userID string) ([]UserProjectBonusDTO, error)
	GetRetroactiveLoggedDays(userID string, startAt string) ([]UserLoggedDaysDTO, error)
}
