package handler

import (
	"dbingo/model"
	dbrepo "dbingo/repo/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := &model.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		db := dbrepo.GetDB()

		ur := dbrepo.NewUserRepository(db)
		err = ur.Create(user)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Only POST method is allowed!")
		return
	}
}
