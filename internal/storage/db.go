package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

func ConnectDB() (*sql.DB, error) {
	conn, err := sql.Open("pgx", buildPgURL())
	if err != nil {
		return nil, fmt.Errorf("failed to create psql connection - %v", err)
	}

	return conn, err
}

func buildPgURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.Get("postgres_user"),
		viper.Get("postgres_password"),
		viper.Get("postgres_hostname"),
		viper.GetInt("postgres_port"),
		viper.Get("postgres_db"),
		viper.Get("postgres_ssl"),
	)
}
