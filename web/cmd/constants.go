package main

type Constants struct {
	DatabaseUrl string
}

func NewConstants() *Constants {
	return &Constants{
		DatabaseUrl: "host=localhost port=5432 dbname=hemnet_gavle user=lindow password=stoffelindow1",
	}
}
