package staff_input

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type EventAttendanceInput struct {
	EventActionName string    `json:"event_action_name"`
	AttendeeLogins  []string  `json:"attendee_logins"`
	PerformedAt     time.Time `json:"performed_at"`
}

type EventAttendanceProcessor struct {
	userRepo       user.UserRepository
	goodActionRepo user.GoodActionRepository
}

func NewEventAttendanceProcessor(userRepo user.UserRepository, goodActionRepo user.GoodActionRepository) *EventAttendanceProcessor {
	return &EventAttendanceProcessor{
		userRepo:       userRepo,
		goodActionRepo: goodActionRepo,
	}
}

func (p *EventAttendanceProcessor) Process(ctx context.Context, input EventAttendanceInput) ([]*user.UserGoodAction, error) {
	eventAction, err := p.goodActionRepo.LoadByName(ctx, input.EventActionName)
	if err != nil {
		return nil, err
	}

	var userActions []*user.UserGoodAction
	var processingErrors []error

	for _, login := range input.AttendeeLogins {
		u, err := p.userRepo.FindByUsername(ctx, login)
		if err != nil {
			processingErrors = append(processingErrors, err)
			continue
		}
		if u == nil {
			processingErrors = append(processingErrors, user.ErrUserNotFound)
			continue
		}

		userGoodAction := user.NewUserGoodAction(
			shared.NewMongoID(),
			u.ID,
			u.Wallet,
			u.Username,
			eventAction.ID,
			eventAction.Name,
			int64(eventAction.RewardAmount),
			input.PerformedAt,
		)

		userActions = append(userActions, userGoodAction)
	}

	if len(processingErrors) > 0 {
		return userActions, fmt.Errorf("some errors occurred while processing event attendance: %v", processingErrors)
	}

	return userActions, nil
}
