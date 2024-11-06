package db

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/ko9motu/go-bun-api/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var (
	dbInstance *bun.DB
	once       sync.Once
)

/*
	fix me !

- マイグレーションコードの作成
*/
func SetupDB() *bun.DB {
	once.Do(func() {
		dsn := "postgres://postgres:postgres@localhost:5432/test_db?sslmode=disable"
		sqldb, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		dbInstance = bun.NewDB(sqldb, pgdialect.New())

		var user models.User

		_, err = dbInstance.NewCreateTable().Model(&user).IfNotExists().Exec(context.Background())
		if err != nil {
			log.Fatalf("Failed to create users table: %v", err)
		}
	})
	return dbInstance
}
