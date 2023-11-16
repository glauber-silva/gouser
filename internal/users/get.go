package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) GetByID(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := Get(h.db, int64(id))
	if err != nil {
		//TODO: validate if the error is due to the user not existing
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)

}

func Get(db *sql.DB, id int64) (*User, error) {
	fmt.Println("Processing Get User")
	stmt := `SELECT * FROM users WHERE id = ?`

	row := db.QueryRow(stmt, id)
	var u User
	fmt.Println("ROW", row)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Age,
		&u.Email,
		&u.Password,
		&u.IsAdmin,
		&u.AddressID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Deleted,
	)

	if err != nil {
		return nil, err
	}

	u.Address, err = GetAddress(db, u.AddressID)
	if err != nil {
		return nil, err
	}

	return &u, nil

}

func GetAddress(db *sql.DB, id int64) (*Address, error) {
	fmt.Println("Processing Get Address")
	fmt.Println("ID", id)
	stmt := `SELECT * FROM addresses WHERE id = ?`
	row := db.QueryRow(stmt, id)
	a := new(Address)
	err := row.Scan(&a.ID, &a.Street, &a.Number, &a.Complement, &a.City, &a.State, &a.Country, &a.PostalCode, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}
