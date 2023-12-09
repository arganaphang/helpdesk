package user

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"

	"github.com/arganaphang/helpdesk/apps/api/domain"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) domain.UserRepository {
	return &repository{db: db}
}

func (r repository) GetByID(ctx context.Context, id ulid.ULID) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	return user, err
}

func (r repository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	return user, err
}

func (r repository) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	if _, err := r.db.NewInsert().Model(&user).Column("name").Column("email").Column("password").Returning("*").Exec(ctx); err != nil {
		return nil, err
	}
	return &user, nil
}
