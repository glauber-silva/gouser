package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

/*
Modify is a handler that receives a JSON with the fields to be updated
And updates the User and Address in the database.
*/
func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Handler Modify")
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	currentUser, err := Get(h.db, int64(userID))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.Address != nil {
		fmt.Println("Current Address ", currentUser.AddressID)
		fmt.Println("Neww Address ", u.Address.ID)
		if currentUser.AddressID != u.Address.ID {
			http.Error(rw, "Address does not belong to the user", http.StatusBadRequest)
			return
		}

		// another way to do it would be deleting the old address and inserting the new one
		fmt.Println("Going to update address", "| u.AddressID: ", u.AddressID, "| u.Address> ", u.Address)
		err = UpdateAddress(h.db, currentUser.AddressID, u.Address)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = Update(h.db, int64(userID), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO: use GET to return the updated user
	// Done, but I have some doubts:
	// Once I using a get above, I would need to use the same code here in case of ORM use
	currentUser, err = Get(h.db, int64(userID))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(currentUser)
}

func Update(db *sql.DB, id int64, u *User) error {
	fmt.Println("UPDATING User", u)
	u.UpdatedAt = time.Now()

	stmt := `UPDATE users SET name = ?, age = ?, email = ?, updated_at = ?, address_id = ? WHERE id = ?`

	_, err := db.Exec(stmt, u.Name, u.Age, u.Email, u.UpdatedAt, u.Address.ID, id)

	return err
}

func UpdateAddress(db *sql.DB, id int64, a *Address) error {
	fmt.Println("UPDATING ADDRESS", a)
	fmt.Println("ID", id)
	currentAddress, err := GetAddress(db, id)
	if err != nil {
		return err
	}

	if a.Street != "" {
		currentAddress.Street = a.Street
	}
	if a.Number != 0 {
		currentAddress.Number = a.Number
	}
	if a.Complement != "" {
		currentAddress.Complement = a.Complement
	}
	if a.City != "" {
		currentAddress.City = a.City
	}
	if a.State != "" {
		currentAddress.State = a.State
	}
	if a.Country != "" {
		currentAddress.Country = a.Country
	}
	if a.PostalCode != "" {
		currentAddress.PostalCode = a.PostalCode
	}

	currentAddress.UpdatedAt = time.Now()

	stmt := `UPDATE addresses SET street = ?, number = ?, complement = ?, city = ?, state = ?, country = ?, postal_code = ?, updated_at = ? WHERE id = ?`
	_, err = db.Exec(stmt,
		currentAddress.Street,
		currentAddress.Number,
		currentAddress.Complement,
		currentAddress.City,
		currentAddress.State,
		currentAddress.Country,
		currentAddress.PostalCode,
		currentAddress.UpdatedAt,
		id)

	return err
}
