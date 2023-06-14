package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	handler "dbingo/http"
	dbrepo "dbingo/repo/db"
)

func main() {
	db := dbrepo.GetDB()
	defer db.Close()

	ur := dbrepo.NewUserRepository(db)
	user, err := ur.GetByID(1)
	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Println("No user with such ID")
	}

	fmt.Println(user)

	users, err := ur.GetAllUsers()
	if err != nil {
		panic(err)
	}

	if users == nil {
		fmt.Println("No users found")
	}

	for _, u := range users {
		fmt.Println(*u)
	}

	http.HandleFunc("/api/users", handler.CreateUser)
	http.ListenAndServe(":3000", nil)
}
