package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/v4/database/pgx" //driver pgx used to run migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
}

func GetConnection(context context.Context) *pgxpool.Pool {
	databaseURL := viper.GetString("database.url")
	conn, err := pgxpool.New(context, "postgres"+databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func RunMigration() {
	databaseURL := viper.GetString("database.url")
	m, err := migrate.New("file://database/migrations", "pgx"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
