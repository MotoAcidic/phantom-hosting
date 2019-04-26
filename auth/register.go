package auth

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	"github.com/jackkdev/phantom-hosting/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var account Account

	db, err := storm.Open("my.db")
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	if len(account.Password) < 6 {
		err := errors.New("Password needs to be more than 6 characters")
		utils.Respond(w, nil, err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	account.ID = 1

	err = db.Save(&account)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, account, nil)
}
