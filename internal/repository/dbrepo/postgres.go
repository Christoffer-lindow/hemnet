package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Christoffer-lindow/portfolio/internal/models"
)

func (m *pgDBRepo) GetProperties(tableName string) ([]models.Property, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var properties []models.Property
	var query string
	if tableName == "gavle" {
		query = `	select
						id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
				from 
					gavle_estates 
				`
	} else if tableName == "stockholm" {
		query = `	select
						id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
					from 
						stockholm_estates 
	`
	} else {
		return properties, errors.New("Could not find tablename")
	}
	rows, err := m.DB.QueryContext(ctx, query)
	for rows.Next() {
		var property models.Property
		err := rows.Scan(
			&property.Id,
			&property.Url,
			&property.ImgUrl,
			&property.Street,
			&property.Address,
			&property.Price,
			&property.Type,
			&property.ContractType,
			&property.Rooms,
			&property.Area,
			&property.Balcony,
			&property.YearBuilt,
			&property.Rent,
			&property.ConsumptionPrice,
			&property.PriceSqm,
		)
		if err != nil {
			log.Println("Got error while scanning", err)
			return properties, err
		}
		boolFalse := false
		if property.Type == "Lägenhet" && property.Balcony == nil {
			property.Balcony = &boolFalse
		}
		properties = append(properties, property)

	}
	if err = rows.Err(); err != nil {
		return properties, nil
	}
	return properties, nil
}

func (m *pgDBRepo) GetSingleProperty(tableName string, id int) (models.Property, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var query string
	var property models.Property
	if tableName == "gavle" {
		query = `select
				id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
			from 
				gavle_estates
			where
				id = $1
				`
	} else if tableName == "stockholm" {
		query = `select
		id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
	from 
		stockholm_estates
	where
		id = $1
		`
	}
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&property.Id,
		&property.Url,
		&property.ImgUrl,
		&property.Street,
		&property.Address,
		&property.Price,
		&property.Type,
		&property.ContractType,
		&property.Rooms,
		&property.Area,
		&property.Balcony,
		&property.YearBuilt,
		&property.Rent,
		&property.ConsumptionPrice,
		&property.PriceSqm,
	)
	if err != nil {
		log.Println("Got error while scanning", err)
		return property, err
	}
	return property, nil
}

func (m *pgDBRepo) GetMostExpensiveProperty(tableName string) (models.Property, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var property models.Property
	var query string
	if tableName == "gavle" {
		query =
			`
	select
		id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
	from 
		gavle_estates
	where price =(select max(price) from gavle_estates)
	group by id
	`
	} else if tableName == "stockholm" {
		query = `
	select
		id, url, img_url, street, address, price, type, contract_type, rooms, area,balcony,year_built,rent, consumption_price, price_sqm
	from 
		stockholm_estates
	where price =(select max(price) from stockholm_estates)
	group by id
	`
	} else {
		return property, errors.New("Could not find tablename")
	}
	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(
		&property.Id,
		&property.Url,
		&property.ImgUrl,
		&property.Street,
		&property.Address,
		&property.Price,
		&property.Type,
		&property.ContractType,
		&property.Rooms,
		&property.Area,
		&property.Balcony,
		&property.YearBuilt,
		&property.Rent,
		&property.ConsumptionPrice,
		&property.PriceSqm,
	)
	if err != nil {
		log.Println("Got error while scanning", err)
		return property, err
	}
	return property, nil
}
