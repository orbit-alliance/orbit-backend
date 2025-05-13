package user

type UserProjectBonusDTO struct {
	ProjectName string
	Points      int
}

type UserLoggedDaysDTO struct {
	Date string
}

type Gateway interface {
	GetRetroativeBonusProject(userID string) ([]UserProjectBonusDTO, error)
	GetRetroativeLoggedDays(userID string, startAt string) ([]UserLoggedDaysDTO, error)
}
