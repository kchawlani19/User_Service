package repository

import (
	"database/sql"
	"errors"
	"go-basic-user-service/model"
)

type DBUserRepository struct {
	db *sql.DB
}

func NewDBUserRepository(db *sql.DB) *DBUserRepository {
	return &DBUserRepository{db: db}
}

func (r *DBUserRepository) Save(user model.User) error {
	_, err := r.db.Exec(
		"Insert INTO users (id,name) VALUE ( ?, ?)",
		user.Id, user.Name,
	)

	return err
}

func (r *DBUserRepository) GetByID(id int) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		"Select id,name from users where id = ?",
		id,
	).Scan(&user.Id, &user.Name)

	if err == sql.ErrNoRows {
		return user, errors.New("user not found")

	}
	return user, err
}

func (r *DBUserRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name = ? WHERE id = ?",
		user.Name, user.Id,
	)
	return err
}

func (r *DBUserRepository) Delete(id int) error {
	_, err := r.db.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)
	return err
}

func (r *DBUserRepository) Exists(id int) (bool, error) {
	var temp int
	err := r.db.QueryRow(
		"SELECT id FROM users WHERE id = ?",
		id,
	).Scan(&temp)

	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}
