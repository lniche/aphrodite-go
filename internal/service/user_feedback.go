package service

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
)

type UserFeedbackService interface {
	GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error)
	AddUserFeedback(ctx context.Context, userCode string, req *v1.AddUserFeedbackRequest) error
}

func NewUserFeedbackService(
	service *Service,
	userFeedbackRepository repository.UserFeedbackRepository,
	userRepo repository.UserRepository,
) UserFeedbackService {
	return &userFeedbackService{
		Service:                service,
		userFeedbackRepository: userFeedbackRepository,
		userRepo:               userRepo,
	}
}

type userFeedbackService struct {
	*Service
	userFeedbackRepository repository.UserFeedbackRepository
	userRepo               repository.UserRepository
}

func (s *userFeedbackService) GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error) {
	return s.userFeedbackRepository.GetUserFeedback(ctx, id)
}

func (s *userFeedbackService) AddUserFeedback(ctx context.Context, userCode string, req *v1.AddUserFeedbackRequest) error {
	_, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	userFeedback := &model.UserFeedback{
		UserCode: userCode,
		Feedback: req.Feedback,
	}
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err = s.userFeedbackRepository.Create(ctx, userFeedback); err != nil {
			return err
		}
		return nil
	})
	return nil
}
