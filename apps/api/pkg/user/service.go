package user

import (
	"context"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/oklog/ulid/v2"
)

type service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) domain.UserService {
	return &service{repo: repo}
}

func (s service) GetByID(ctx context.Context, id ulid.ULID) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s service) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s service) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	return s.repo.Create(ctx, user)
}
