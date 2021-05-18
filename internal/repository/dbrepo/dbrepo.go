package dbrepo

import (
	"database/sql"

	"github.com/Christoffer-lindow/portfolio/internal/config"
	"github.com/Christoffer-lindow/portfolio/internal/repository"
)

type pgDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPgRepo(conn *sql.DB, a *config.AppConfig) repository.DBRepo {
	return &pgDBRepo{
		App: a,
		DB:  conn,
	}
}
