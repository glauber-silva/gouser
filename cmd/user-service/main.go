package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/glauber-silva/gouser/internal/health"
	"github.com/glauber-silva/gouser/internal/users"
	db "github.com/glauber-silva/gouser/pkg/database"
	"github.com/go-chi/chi"
)

func main() {
	db := getSessions()

	r := chi.NewRouter()

	// Set routes
	health.SetRoutes(r, db)
	users.SetRoutes(r, db)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
	log.Println("Listing requests at http://localhost:8080/health")

}

func getSessions() *sql.DB {
	db, err := db.OpenConnection()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
