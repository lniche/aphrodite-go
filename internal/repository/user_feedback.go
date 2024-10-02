package repository

import (
	"aphrodite-go/internal/model"
	"context"
)

type UserFeedbackRepository interface {
	GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error)
	Create(ctx context.Context, userFeedback *model.UserFeedback) error
}

func NewUserFeedbackRepository(
	r *Repository,
) UserFeedbackRepository {
	return &userFeedbackRepository{
		Repository: r,
	}
}

type userFeedbackRepository struct {
	*Repository
}

func (r *userFeedbackRepository) GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error) {
	var userFeedback model.UserFeedback

	return &userFeedback, nil
}

func (r *userFeedbackRepository) Create(ctx context.Context, userFeedback *model.UserFeedback) error {
	if err := r.DB(ctx).Create(userFeedback).Error; err != nil {
		return err
	}
	return nil
}
