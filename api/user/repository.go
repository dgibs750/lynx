package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UserById(id uuid.UUID) (*User, error) {
	user := &User{}
	row := r.db.QueryRow("SELECT * FROM user WHERE id = UUID_TO_BIN(?)", id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MasterPassword, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UserByEmail(email string) (*User, error) {
	user := &User{}
	row := r.db.QueryRow("SELECT * FROM user WHERE email = ?", email)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MasterPassword, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) AddUser(newUser *NewUser) (*User, error) {
	user := &User{}
	_, err := r.db.Exec(
		"INSERT INTO user (first_name, last_name, email, master_password) VALUES (?,?,?,?)",
		newUser.FirstName, newUser.LastName, newUser.Email, newUser.MasterPassword,
	)
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow("SELECT * FROM user WHERE email = ?", newUser.Email)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MasterPassword, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(updateUserData *UpdateUserData) (*User, error) {
	user := &User{}
	_, err := r.db.Exec(
		"UPDATE user SET first_name = ?, last_name = ?, email = ?, master_password = ? WHERE id = UUID_TO_BIN(?)",
		updateUserData.FirstName, updateUserData.LastName, updateUserData.Email, updateUserData.MasterPassword, updateUserData.ID,
	)
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow("SELECT * FROM user WHERE id = UUID_TO_BIN(?)", updateUserData.ID)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MasterPassword, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}
