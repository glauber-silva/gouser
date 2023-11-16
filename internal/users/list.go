package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) List(rw http.ResponseWriter, r *http.Request) {
	users, err := SelectAll(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(users)

}

func SelectAll(db *sql.DB) ([]User, error) {
	stmt := `SELECT * FROM users where deleted = false`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User

		err := rows.Scan(
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
		users = append(users, u)
	}

	return users, nil

}
