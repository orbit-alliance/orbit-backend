package user

type Gateway interface {
	GetRetroativeBonusProject(userID string) (map[string]int)
	GetRetroativeLoggedDays(userID string, startAt string) (map[string]string)
}
