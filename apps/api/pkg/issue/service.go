package issue

import (
	"context"

	"github.com/google/uuid"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/arganaphang/helpdesk/apps/api/dto"
)

type service struct {
	repo domain.IssueRepository
}

func NewService(repo domain.IssueRepository) domain.IssueService {
	return &service{repo: repo}
}

func (s service) Create(ctx context.Context, issue domain.Issue) (*domain.Issue, error) {
	return s.repo.Create(ctx, issue)
}

func (s service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Issue, error) {
	return s.repo.GetByID(ctx, id)
}

func (s service) Take(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Issue, error) {
	return s.repo.Take(ctx, id, userID)
}

func (s service) Solve(ctx context.Context, id uuid.UUID) (*domain.Issue, error) {
	return s.repo.Solve(ctx, id)
}

func (s service) Get(ctx context.Context, queryParams dto.IssueQueryParams) ([]domain.Issue, *int, error) {
	return s.repo.Get(ctx, queryParams)
}
