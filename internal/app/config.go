package app

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

type PgConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func GetPort() string {
	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))
	if len(port) == 0 {
		logrus.Fatalf("invalid port: %s", port)
	}
	return port
}

func GetPgConfig() string {
	pg := PgConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
	if len(pg.User) == 0 {
		logrus.Warn("ivalid PGUSER")
	}
	if len(pg.Password) == 0 {
		logrus.Warn("ivalid PGPASSWORD")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg.User, pg.Password, pg.Host, pg.Port, pg.Database)
}

func PrepareDBPool(c *Config) *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), c.PgConnStr)
	if err != nil {
		logrus.Warnf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		logrus.Info("connected to postgres")
	}
	return dbPool
}
