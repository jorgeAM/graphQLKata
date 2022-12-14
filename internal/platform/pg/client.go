package pg

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Client struct {
	*bun.DB
}

func NewClient() (*Client, error) {
	dsn := "postgres://admin:123456@localhost:5432/graphqlgo?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.Ping()

	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}
