package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type BonusProjectService struct {
	userRepo           user.UserRepository
	eventBus           *shared.EventBus
	userProjectRepo    user.UserProjectRepository
	userGoodActionRepo user.UserGoodActionRepository
	goodActionRepo     user.GoodActionRepository
	api42Gateway       user.Api42Gateway
}

func NewBonusProjectService(
	userRepo user.UserRepository,
	eventBus *shared.EventBus,
	userProjectRepo user.UserProjectRepository,
	userGoodActionRepo user.UserGoodActionRepository,
	goodActionRepo user.GoodActionRepository,
	api42Gateway user.Api42Gateway,
) *BonusProjectService {
	return &BonusProjectService{
		userRepo:           userRepo,
		eventBus:           eventBus,
		userProjectRepo:    userProjectRepo,
		userGoodActionRepo: userGoodActionRepo,
		goodActionRepo:     goodActionRepo,
		api42Gateway:       api42Gateway,
	}
}

func (s *BonusProjectService) ApplyRetroactiveBonusProject(ctx context.Context, event shared.DomainEvent) {

	fmt.Println("Applying retroactive bonus project...")

	evt, ok := event.(*user.User42Registered)
	if !ok {
		return
	}

	userID := shared.ObjectIDToString(evt.User.ID)

	usr, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		fmt.Printf("Error finding user with ID %s: %v\n", userID, err)
		return
	}

	bonusProjects, err := s.api42Gateway.GetRetroactiveBonusProject(usr.ID42)
	if err != nil {
		fmt.Printf("Error fetching retroactive bonus projects for user %s: %v\n", usr.ID42, err)
		return
	}

	didProjects, err := s.userProjectRepo.LoadByUserID(ctx, shared.ObjectIDToString(usr.ID))
	if err != nil {
		fmt.Printf("Error loading user projects for user %s: %v\n", usr.ID42, err)
		return
	}

	goodAction, err := s.goodActionRepo.LoadByName(ctx, "Bônus de Projeto")
	if err != nil {
		fmt.Printf("Error loading good action 'Bônus de Projeto': %v\n", err)
		return
	}

	err = s.processBonusProjects(ctx, usr, bonusProjects, didProjects, goodAction)
	if err != nil {
		fmt.Printf("Error processing bonus projects for user %s: %v\n", usr.ID42, err)
		return
	}
}

func (s *BonusProjectService) processBonusProjects(ctx context.Context, usr *user.User, bonusProjects []user.UserProjectBonusDTO, didProjects []*user.UserProject, goodAction *coin.GoodAction) error {
	for _, bp := range bonusProjects {
		if err := s.applySingleBonus(ctx, usr, bp, didProjects, goodAction); err != nil {
			return err
		}
	}
	return nil
}

func (s *BonusProjectService) applySingleBonus(ctx context.Context, usr *user.User, bp user.UserProjectBonusDTO, didProjects []*user.UserProject, goodAction *coin.GoodAction) error {
	oldProject := s.findSameProject(bp, didProjects)
	newProject := user.NewUserProject(shared.NewMongoID(), bp.ProjectName, uint8(bp.Points), usr.ID, usr.Username)

	if oldProject == nil || oldProject.MaxScore < newProject.MaxScore {
		if err := s.userProjectRepo.Save(ctx, newProject); err != nil {
			return err
		}
	}

	evt, userGoodAction, err := usr.DoBonusProject(newProject, oldProject, *goodAction)
	if err != nil {
		return err
	}

	if evt != nil {
		s.eventBus.Publish(evt)

		if err := s.userRepo.Save(ctx, usr); err != nil {
			return err
		}
		if err := s.userGoodActionRepo.Save(ctx, userGoodAction); err != nil {
			fmt.Printf("Error saving user good action: %v\n", err)
			return err
		}
	}
	return nil
}

func (s *BonusProjectService) findSameProject(bonusProject user.UserProjectBonusDTO, projects []*user.UserProject) *user.UserProject {
	for _, project := range projects {
		if project.Name == bonusProject.ProjectName {
			return project
		}
	}
	return nil
}

func (s *BonusProjectService) DailyBonusProjectJob() {
	ctx := context.Background()

	users, err := s.userRepo.LoadAll(ctx)
	if err != nil {
		fmt.Printf("Erro ao carregar usuários: %v\n", err)
		return
	}

	const maxConcurrency = 10
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for _, u := range users {
		sem <- struct{}{} // ocupa uma "vaga" do semáforo
		wg.Add(1)

		go func(u *user.User) {
			defer func() {
				<-sem     // libera a vaga
				wg.Done() // sinaliza fim da goroutine
			}()

			s.ApplyRetroactiveBonusProject(ctx, &user.User42Registered{
				User: u, // importante: desreferenciar o ponteiro aqui
			})
		}(u)
	}

	wg.Wait() // aguarda todas as goroutines finalizarem
}
