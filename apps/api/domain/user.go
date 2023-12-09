package domain

import (
	"context"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:todos"`
	ID            ulid.ULID  `bun:"id,pk"       json:"id"`
	Name          string     `bun:"name"        json:"name"`
	Email         string     `bun:"email"       json:"email"`
	Password      string     `bun:"password"    json:"-"`
	IsValid       bool       `bun:"is_valid"    json:"-"`
	CreatedAt     time.Time  `bun:"created_at"  json:"-"`
	UpdatedAt     *time.Time `bun:"updated_at"  json:"-"`
}

type UserService interface {
	GetByID(ctx context.Context, id ulid.ULID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id ulid.ULID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}
