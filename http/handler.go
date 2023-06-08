package handler

import (
	"dbingo/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := &model.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(user)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		fmt.Println("Only POST method is allowed!")
		return
	}
}
