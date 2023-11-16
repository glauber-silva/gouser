package users

import (
	"database/sql"

	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.Create)

		r.Group(func(r chi.Router) {
			r.Put("/{id}", h.Modify)
			r.Delete("/{id}", h.Delete)
			r.Get("/{id}", h.GetByID)
			r.Get("/", h.List)
		})

	})

}
