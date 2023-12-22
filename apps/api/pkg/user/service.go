package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/arganaphang/helpdesk/apps/api/domain"
)

type service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) domain.UserService {
	return &service{repo: repo}
}

func (s service) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s service) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s service) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	return s.repo.Create(ctx, user)
}
