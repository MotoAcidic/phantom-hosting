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

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/generateconfigfile", GenerateConfigFile).Methods("POST")
	r.HandleFunc("/generatemasternodenstring", GenerateMasternodeString).Methods("POST")

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GenerateConfigFile(w http.ResponseWriter, r *http.Request) {
	config.GenerateConfigurationFile("masternode.txt")
}

func GenerateMasternodeString(w http.ResponseWriter, r *http.Request) {
	var mnConfig config.MasternodeString

	err := json.NewDecoder(r.Body).Decode(&mnConfig)

	if err != nil {
		log.Println(err.Error())
	}

	if mnConfig.TransactionID == "" {
		fmt.Println("A TxID is needed")
		utils.Respond(w, nil, err)
		return
	}

	mnString := config.GenerateNodeDetails(mnConfig)

	utils.Respond(w, mnString, nil)
}