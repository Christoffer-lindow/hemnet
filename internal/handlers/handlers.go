package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Christoffer-lindow/portfolio/internal/config"
	"github.com/Christoffer-lindow/portfolio/internal/driver"
	"github.com/Christoffer-lindow/portfolio/internal/repository"
	"github.com/Christoffer-lindow/portfolio/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	Db  repository.DBRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		Db:  dbrepo.NewPgRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Properties(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	city := strings.SplitAfter(r.RequestURI, "/properties/")[1]

	properties, err := m.Db.GetProperties(city)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not get properties from " + city})
		return
	}
	json.NewEncoder(w).Encode(properties)
}

func (m *Repository) SingleProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	tmp := strings.SplitAfter(r.RequestURI, "/properties/single/")[1]
	id := strings.SplitAfter(tmp, "/")[0]
	id = strings.Replace(id, "/", "", 1)
	log.Println(id)
	idVal, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Could not parse id to integer")
	}
	city := strings.SplitAfter(r.RequestURI, "/city/")[1]
	property, err := m.Db.GetSingleProperty(city, idVal)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not find property"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(property)
}

func (m *Repository) MostExpensiveProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")
	city := strings.SplitAfter(r.RequestURI, "/max_price/")[1]
	property, err := m.Db.GetMostExpensiveProperty(city)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not find property"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(property)
}
