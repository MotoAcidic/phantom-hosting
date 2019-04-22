package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/jackkdev/phantom-hosting-api/config"
	"gitlab.com/jackkdev/phantom-hosting-api/utils"
	"log"
	"net/http"
)

var mnString string

func Start() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generateconfigfile", GenerateConfigFile).Methods("POST")
	api.HandleFunc("/generatemasternodestring", GenerateMasternodeString).Methods("POST")
	api.HandleFunc("/addmasternode", AddMasternode).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// API Handlers
func GenerateConfigFile(w http.ResponseWriter, r *http.Request) {
	config.GenerateConfigurationFile("masternode.txt")
	utils.Respond(w, "Configuration file created", nil)
}

func GenerateMasternodeString(w http.ResponseWriter, r *http.Request) {
	var mnConfig config.MasternodeString

	err := json.NewDecoder(r.Body).Decode(&mnConfig)

	if err != nil {
		log.Println(err.Error())
	}

	mnString, err = config.GenerateNodeDetails(mnConfig)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, mnString, nil)
}

func AddMasternode(w http.ResponseWriter, r *http.Request) {
	config.AddMasternodeToConfigFile("masternode.txt", mnString)
	utils.Respond(w, "Masternode added successfully to configuration file", nil)
}