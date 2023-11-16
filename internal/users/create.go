package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	var u User

	fmt.Println("PRINTANDO: ", r.Body)

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	addressId, err := InsertAddress(h.db, &u.Address)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("addressId: ", addressId)

	u.AddressID = addressId
	userid, err := Insert(h.db, &u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("userid: ", userid)
	fmt.Println("u: ", u)
	u.ID = userid
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&u)

}

func Insert(db *sql.DB, u *User) (int64, error) {
	stmt := `INSERT INTO users (name, age, email, password, is_admin, address_id)
	VALUES (?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(stmt, u.Name, u.Age, u.Email, u.Password, u.IsAdmin, u.AddressID)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil

}

func InsertAddress(db *sql.DB, a *Address) (int64, error) {
	stmt := `INSERT INTO addresses (street, number, complement, city, state, country, postal_code)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(stmt, a.Street, a.Number, a.Complement, a.City, a.State, a.Country, a.PostalCode)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}
	return id, nil
}
