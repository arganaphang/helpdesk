package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/arganaphang/helpdesk/apps/api/dto"
)

type Issue struct {
	bun.BaseModel `bun:"table:issues"`
	ID            uuid.UUID  `bun:"id,pk"          json:"id"`
	Title         string     `bun:"title"          json:"title"`
	Detail        string     `bun:"detail"         json:"detail"`
	CustomerName  string     `bun:"customer_name"  json:"customer_name"`
	CustomerEmail string     `bun:"customer_email" json:"customer_email"`
	TakenBy       *uuid.UUID `bun:"taken_by"       json:"taken_by"`
	SolvedAt      *time.Time `bun:"solved_at"      json:"solved_at"`
	CreatedAt     time.Time  `bun:"created_at"     json:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"     json:"-"`
}

type IssueService interface {
	Create(ctx context.Context, issue Issue) (*Issue, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Issue, error)
	Take(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*Issue, error)
	Solve(ctx context.Context, id uuid.UUID) (*Issue, error)
	Get(ctx context.Context, queryParams dto.IssueQueryParams) ([]Issue, *int, error)
}

type IssueRepository interface {
	Create(ctx context.Context, issue Issue) (*Issue, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Issue, error)
	Take(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*Issue, error)
	Solve(ctx context.Context, id uuid.UUID) (*Issue, error)
	Get(ctx context.Context, queryParams dto.IssueQueryParams) ([]Issue, *int, error)
}
