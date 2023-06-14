package db

import (
	"database/sql"
	"dbingo/model"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByID(id int64) (*model.User, error) {
	stmt, err := u.db.Prepare("SELECT id, first_name, last_name FROM users WHERE users.id = $1")
	if err != nil {
		return nil, errors.New("can not prepare statement to get users by id")
	}

	row := stmt.QueryRow(id)
	err = row.Err()
	if err != nil {
		return nil, errors.New("select user failed")
	}

	user := &model.User{}
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.New("types mismatch during the scanning")
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() ([]*model.User, error) {
	stmt, err := u.db.Prepare("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, errors.New("can not prepare statement to get all users")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("select user failed")
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) Create(user *model.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users (first_name, middle_name, last_name) VALUES ($1, $2, $3)")
	if err != nil {
		return errors.New("can not prepare statement to insert new user")
	}

	_, err = stmt.Exec(user.FirstName, user.MiddleName, user.LastName)
	if err != nil {
		return errors.New("insert into users failed")
	}

	return nil
}
