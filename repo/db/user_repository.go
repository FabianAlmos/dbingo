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
	stmt, err := u.db.Prepare("SELECT id, first_name || ' ' || last_name FROM person WHERE person.id = $1")
	if err != nil {
		return nil, errors.New("can not prepare statement to get users by id")
	}

	row := stmt.QueryRow(id)
	err = row.Err()
	if err != nil {
		return nil, errors.New("select user failed")
	}

	user := &model.User{}
	err = row.Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.New("types mismatch during the scanning")
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers(limit, offset uint) ([]*model.User, error) {
	stmt, err := u.db.Prepare("SELECT id, concat(first_name, ' ', last_name) FROM person LIMIT $1 from $2")
	if err != nil {
		return nil, errors.New("can not prepare statement to get all user")
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, errors.New("select user failed")
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Name)
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
	stmt, err := u.db.Prepare("INSERT INTO person (id, first_name, middle_name, last_name, nickname) VALUES (id, first_name, middle_name, last_name, nickname)")
	if err != nil {
		return errors.New("can not prepare statement to get all user")
	}
	row := stmt.QueryRow(user.ID, user.Name)
	err = row.Err()
	if err != nil {
		return errors.New("insert into users failed")
	}
	return nil
}
