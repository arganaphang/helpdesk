package domain

import (
	"context"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

type Issue struct {
	bun.BaseModel `bun:"table:todos"`
	ID            ulid.ULID  `bun:"id,pk"          json:"id"`
	Title         string     `bun:"title"          json:"title"`
	Detail        string     `bun:"detail"         json:"detail"`
	CustomerName  string     `bun:"customer_name"  json:"customer_name"`
	CustomerEmail string     `bun:"customer_email" json:"customer_email"`
	TakenBy       *ulid.ULID `bun:"taken_by"       json:"taken_by"`
	SolvedAt      *time.Time `bun:"solved_at"      json:"solved_at"`
	CreatedAt     time.Time  `bun:"created_at"     json:"-"`
	UpdatedAt     *time.Time `bun:"updated_at"     json:"-"`
}

type IssueService interface {
	Create(ctx context.Context, issue Issue) (*Issue, error)
	GetByID(ctx context.Context, id ulid.ULID) (*Issue, error)
	Get(ctx context.Context) ([]Issue, error)
}

type IssueRepository interface {
	Create(ctx context.Context, issue Issue) (*Issue, error)
	GetByID(ctx context.Context, id ulid.ULID) (*Issue, error)
	Get(ctx context.Context) ([]Issue, error)
}
