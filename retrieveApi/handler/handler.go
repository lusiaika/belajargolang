package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"retrieveApi/entity"
)

func DataUser(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10.")
	if err != nil {
		log.Fatalln(err)
	}

	var result []entity.ResUser

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	var users []entity.User
	var user entity.User
	for _, data := range result {
		user = entity.User{
			Id:         data.Id,
			Uid:        data.Uid,
			First_name: data.First_name,
			Last_name:  data.Last_name,
			Username:   data.Username,
			Address:    data.Address,
		}
		users = append(users, user)
	}

	writeJsonResp(w, statusSuccess, users)

}
