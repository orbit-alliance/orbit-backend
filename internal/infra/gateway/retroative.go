package gateway_42

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func GetRetroativeBonusProject(userID string) ([]user.UserProjectBonusDTO, error) {
	token, err := GetToken()

	bonusList := []user.UserProjectBonusDTO{}
	apiData, err := GetUserByID(userID, token)

	if err != nil {
		return nil, ErrFailToGetUserIn42
	}

	for _, project := range apiData.ProjectsUsers {
		if project.FinalMark > 100 {
			newProject := user.UserProjectBonusDTO{
				ProjectName: project.Project.Name,
				Points:      project.FinalMark,
			}
			bonusList = append(bonusList, newProject)
		}
	}
	return bonusList, nil
}

func GetRetroativeLoggedDays(userID string, startAt string) ([]user.UserLoggedDaysDTO, error) {
	token, err := GetToken()
	endAt := time.Now().UTC().Format(time.RFC3339Nano)
	dayList := []user.UserLoggedDaysDTO{}

	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	apiData, err := GetLocationByUserID(startAt, endAt, userID, token)

	if err != nil {
		return nil, ErrFailToGetLocationIn42
	}

	for date, _ := range apiData {
		newDay := user.UserLoggedDaysDTO{
			Date: date,
		}
		dayList = append(dayList, newDay)
	}
	return dayList, nil
}
