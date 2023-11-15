package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//err = u.Validate()
	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = Update(h.db, int64(userID), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO: use GET to return the updated user
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Update(db *sql.DB, id int64, u *User) error {

	u.UpdatedAt = time.Now()

	stmt := `UPDATE users SET name = $1, age = $2, email = $3, is_admin = $4, updated_at = $5 address_id = $6 WHERE id = $7`

	_, err := db.Exec(stmt, u.Name, u.Age, u.Email, u.IsAdmin, u.UpdatedAt, u.AddressID, id)

	return err

}
