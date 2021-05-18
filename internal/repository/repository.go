package repository

import "github.com/Christoffer-lindow/portfolio/internal/models"

type DBRepo interface {
	GetProperties(tableName string) ([]models.Property, error)
	GetSingleProperty(tableName string, id int) (models.Property, error)
	GetMostExpensiveProperty(tableName string) (models.Property, error)
}
