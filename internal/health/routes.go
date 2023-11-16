package health

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {

	h := handler{db}

	r.Route("/health", func(r chi.Router) {
		r.Get("/", h.HealthCheck)

	})

}

func (h *handler) HealthCheck(rw http.ResponseWriter, r *http.Request) {
	health, err := CheckDB(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]bool{"db": health})

}

func CheckDB(db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		return false, err
	}

	return true, nil

}
