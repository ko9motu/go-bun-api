package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/ko9motu/go-bun-api/models"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var Db *bun.DB

/*
	fix me !

- マイグレーションコードの作成
*/
func init() {
	dsn := "postgres://postgres:postgres@db:5432/test_db?sslmode=disable"
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Db = bun.NewDB(sqldb, pgdialect.New())

	var user models.User

	_, err = Db.NewCreateTable().Model(&user).IfNotExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}
