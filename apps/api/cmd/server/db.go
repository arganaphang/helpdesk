package main

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"
)

func NewDatabase() *bun.DB {
	db := bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL")))),
		pgdialect.New(),
	)
	hook := bundebug.NewQueryHook()
	if os.Getenv("MODE") == "development" {
		hook = bundebug.NewQueryHook(bundebug.WithVerbose(true))
	}
	db.AddQueryHook(hook)

	if err := db.Ping(); err != nil {
		zap.L().Fatal("Failed to connect to database")
	}
	return db
}
