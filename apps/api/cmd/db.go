package main

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

func NewDatabase() *bun.DB {
	db := bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL")))),
		pgdialect.New(),
	)
	if err := db.Ping(); err != nil {
		zap.L().Fatal("Failed to connect to database")
	}
	return db
}
