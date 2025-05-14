package gateway_42

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (r *Gateway42) GetRetroativeBonusProject(userID string) ([]user.UserProjectBonusDTO, error) {
	token, err := getToken()
	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	bonusList := []user.UserProjectBonusDTO{}
	apiData, err := getUserByID(userID, token)

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

func (r *Gateway42) GetRetroativeLoggedDays(userID string, startAt string) ([]user.UserLoggedDaysDTO, error) {
	token, err := getToken()
	endAt := time.Now().UTC().Format(time.RFC3339Nano)
	dayList := []user.UserLoggedDaysDTO{}

	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	apiData, err := getLocationByUserID(startAt, endAt, userID, token)

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
