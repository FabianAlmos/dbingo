package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	handler "dbingo/http"
	dbrepo "dbingo/repo/db"
)

func main() {
	connStr := "postgres://dbdata:dbdatapswd@localhost/lesson"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ur := dbrepo.NewUserRepository(db)
	user, err := ur.GetByID(0)
	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Println("No user with such ID")
	}

	fmt.Println(user)

	page := 2
	pageSize := 1

	var (
		limit  uint = 1
		offset uint = 0
	)

	users, err := ur.GetAllUsers(limit, offset)
	if err != nil {
		panic(err)
	}

	if users == nil {
		fmt.Println("No users found")
	}

	for _, u := range users {
		fmt.Println(*u)
	}

	http.HandleFunc("/api/users", handler.PostUser)
	http.ListenAndServe(":3000", nil)
}
