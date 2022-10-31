package store

import (
	"strconv"

	"github.com/goocarry/cnord-test/internal/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(firstname string, lastname string) (string, error) {
	var id string

	err := r.store.db.QueryRow(`WITH e AS (INSERT INTO public."users" (user_id, firstname, lastname) 
		VALUES (DEFAULT, $1, $2)
		ON CONFLICT("firstname", "lastname") DO NOTHING
    	RETURNING user_id)
		SELECT * FROM e
		UNION SELECT user_id FROM public."users" WHERE firstname=$1 AND lastname=$2`,
		firstname,
		lastname,
	).Scan(&id)
	if err != nil {
		r.store.log.Printf("error-25c065c1: cannot create user, error: %v", err)
		return "", err
	}
	return id, nil
}

// GetByID ...
func (r *UserRepository) GetByID(ID string) (*model.User, error) {

	var user model.User

	intID, err := strconv.Atoi(ID)
	if err != nil {
		r.store.log.Println("cant convert")
	}

	err = r.store.db.QueryRow(`SELECT firstname, lastname FROM public."users" WHERE user_id=$1`,
		intID,
	).Scan(&user.FirstName, &user.LastName)
	if err != nil {
		r.store.log.Printf("error-25c065c1: cannot get user, error: %v", err)
		return nil, err
	}
	return &user, nil
}

