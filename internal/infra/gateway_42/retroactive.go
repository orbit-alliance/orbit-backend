package gateway_42

import (
	"context"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (g *Gateway42) GetRetroactiveBonusProject(user42ID string) ([]user.UserProjectBonusDTO, error) {
	ctx := context.Background()
	token, err := g.getToken(ctx)
	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	bonusList := []user.UserProjectBonusDTO{}
	apiData, err := g.getUserByID(ctx, user42ID, token)

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

func (g *Gateway42) GetRetroactiveLoggedDays(
	user42ID string,
	startAt string,
) ([]user.UserLoggedDaysDTO, error) {

	ctx := context.Background()
	token, err := g.getToken(ctx)
	if err != nil {
		return nil, ErrFailToCreate42Token
	}

	endAt := time.Now().UTC().Format(time.RFC3339Nano)

	apiData, err := g.getLocationByUserID(
		ctx,
		user42ID,
		startAt,
		endAt,
		token,
	)
	if err != nil {
		return nil, ErrFailToGetLocationIn42
	}

	dayList := make([]user.UserLoggedDaysDTO, 0, len(apiData))
	for date := range apiData {
		dayList = append(dayList, user.UserLoggedDaysDTO{Date: date})
	}
	return dayList, nil
}
