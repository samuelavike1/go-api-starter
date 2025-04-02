package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang/config"
)

func ConnectDB(cfg config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return conn, nil
}
