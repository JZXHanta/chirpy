package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JZXHanta/chirpy/internal/auth"
	"github.com/JZXHanta/chirpy/internal/database"
)

type User struct {
<<<<<<< HEAD
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
=======
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}

	user, err := cfg.DB.CreateUser(params.Email, hashedPassword)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			respondWithError(w, http.StatusConflict, "User already exists")
			return
		}

		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		User: User{
<<<<<<< HEAD
			ID:          user.ID,
			Email:       user.Email,
			IsChirpyRed: user.IsChirpyRed,
=======
			ID:    user.ID,
			Email: user.Email,
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
		},
	})
}
