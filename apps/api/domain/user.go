package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            uuid.UUID  `bun:"id,pk"       json:"id"`
	Name          string     `bun:"name"        json:"name"`
	Email         string     `bun:"email"       json:"email"`
	Password      string     `bun:"password"    json:"-"`
	IsValid       bool       `bun:"is_valid"    json:"-"`
	CreatedAt     time.Time  `bun:"created_at"  json:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"  json:"-"`
}

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}
