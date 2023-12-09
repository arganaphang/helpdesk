package issue

import (
	"context"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/oklog/ulid/v2"
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

func (s service) GetByID(ctx context.Context, id ulid.ULID) (*domain.Issue, error) {
	return s.repo.GetByID(ctx, id)
}

func (s service) Get(ctx context.Context) ([]domain.Issue, error) {
	return s.repo.Get(ctx)
}
