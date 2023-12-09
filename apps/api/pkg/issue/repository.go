package issue

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"

	"github.com/arganaphang/helpdesk/apps/api/domain"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) domain.IssueRepository {
	return &repository{db: db}
}

func (r repository) Create(ctx context.Context, issue domain.Issue) (*domain.Issue, error) {
	if _, err := r.db.NewInsert().Model(&issue).Column("title").Column("detail").Column("customer_name").Column("customer_email").Returning("*").Exec(ctx); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (r repository) GetByID(ctx context.Context, id ulid.ULID) (*domain.Issue, error) {
	issue := &domain.Issue{}
	err := r.db.NewSelect().Model(issue).Where("id = ?", id).Scan(ctx)
	return issue, err
}

func (r repository) Get(ctx context.Context) ([]domain.Issue, error) {
	issues := []domain.Issue{}
	err := r.db.NewSelect().Model(issues).Order("created_at asc").Scan(ctx)
	return issues, err
}
