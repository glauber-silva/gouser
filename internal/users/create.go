package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//err = u.Validate()
	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	id, err := Insert(h.db, &u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = id
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&u)

}

func Insert(db *sql.DB, u *User) (int64, error) {
	stmt := `INSERT INTO users (name, age, email, password, is_admin, address_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(stmt, u.Name, u.Age, u.Email, u.Password, u.IsAdmin, u.AddressID)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()

}
