package users

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = Delete(h.db, int64(userID))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")

}

func Delete(db *sql.DB, id int64) error {
	now := time.Now()

	stmt := `UPDATE users SET updated_at = $1 deleted = $2 WHERE id = $3`
	_, err := db.Exec(stmt, now, true, id)

	return err
}
