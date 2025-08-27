package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getusersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createusersHandler(w http.ResponseWriter, r *http.Request) {
	//decode request body to Useer struct

	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	//input validation
	if u.FirstName == "" {
		return errors.New("First name is required")
	}
	if u.LastName == "" {
		return errors.New("Last name is required")
	}

	//storage validation

	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User aready exists")
		}
	}

	users = append(users, u)
	return nil
}
