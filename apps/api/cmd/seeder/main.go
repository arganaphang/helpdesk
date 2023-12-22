package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	faker "github.com/go-faker/faker/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"

	"github.com/arganaphang/helpdesk/apps/api/domain"
)

func newDatabase() *bun.DB {
	db := bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL")))),
		pgdialect.New(),
	)
	if err := db.Ping(); err != nil {
		zap.L().Fatal("Failed to connect to database")
	}
	return db
}

func main() {
	data := []domain.Issue{}
	n := 100_000
	db := newDatabase()
	ctx := context.Background()

	for i := 0; i < n; i++ {
		data = append(data, domain.Issue{
			Title:         faker.Sentence(),
			Detail:        faker.Paragraph(),
			CustomerName:  faker.Name(),
			CustomerEmail: faker.Email(),
		})
	}
	if _, err := db.NewInsert().Model(&data).Column("title").Column("detail").Column("customer_name").Column("customer_email").Exec(ctx); err != nil {
		log.Println("Failed to seed", err.Error())
		return
	}
	log.Println("Seeding success")
}
