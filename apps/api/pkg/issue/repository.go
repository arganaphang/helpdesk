package issue

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/arganaphang/helpdesk/apps/api/dto"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) domain.IssueRepository {
	return &repository{db: db}
}

func (r repository) Create(ctx context.Context, issue domain.Issue) (*domain.Issue, error) {
	_, err := r.db.NewInsert().
		Model(&issue).
		Column("title").
		Column("detail").
		Column("customer_name").
		Column("customer_email").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

func (r repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Issue, error) {
	issue := &domain.Issue{}
	err := r.db.NewSelect().Model(issue).Where("id = ?", id).Scan(ctx)
	return issue, err
}

func (r repository) Take(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Issue, error) {
	issue := &domain.Issue{}
	_, err := r.db.NewUpdate().Model(issue).Set("taken_by = ?", userID).Where("id = ?", id).Returning("*").Exec(ctx)
	return issue, err
}

func (r repository) Solve(ctx context.Context, id uuid.UUID) (*domain.Issue, error) {
	issue := &domain.Issue{}
	_, err := r.db.NewUpdate().
		Model(issue).
		Set("solved_at = ?", time.Now()).
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	return issue, err
}

func (r repository) Get(ctx context.Context, queryParams dto.IssueQueryParams) ([]domain.Issue, *int, error) {
	issues := make([]domain.Issue, 0)
	query := r.db.NewSelect().Model((*domain.Issue)(nil))

	if queryParams.Q != nil {
		query.Where("title ILIKE ?", "%"+*queryParams.Q+"%")
	}

	if queryParams.CustomerName != nil {
		query.Where("customer_name ILIKE ?", "%"+*queryParams.CustomerName+"%")
	}

	if queryParams.CustomerEmail != nil {
		query.Where("customer_email ILIKE ?", "%"+*queryParams.CustomerEmail+"%")
	}

	if queryParams.TakenBy != nil {
		query.Where("taken_by = ?", queryParams.TakenBy.String())
	}

	if queryParams.IsSolved != nil {
		if *queryParams.IsSolved {
			query.Where("solved_at IS NOT NULL")
		} else {
			query.Where("solved_at IS NULL")
		}
	}
	sortBy := "id"
	if queryParams.SortBy != nil {
		sortBy = *queryParams.SortBy
	}

	sortType := "asc"
	if queryParams.SortType != nil {
		sortType = *queryParams.SortType
	}

	count, err := query.Order(fmt.Sprintf("%s %s", sortBy, sortType)).
		Limit(queryParams.Limit).
		Offset(queryParams.Offset).
		ScanAndCount(ctx, &issues)
	if err != nil {
		return nil, nil, err
	}

	return issues, &count, err
}
