package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/jackkdev/phantom-hosting-api/config"
	"gitlab.com/jackkdev/phantom-hosting-api/utils"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/generateconfig", GenerateConfig).Methods("GET")

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GenerateConfig(w http.ResponseWriter, r *http.Request) {
	configString := config.GenerateNodeDetails(config.MasternodeString{})
	utils.Respond(w, configString, nil)
}