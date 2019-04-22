package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/jackkdev/phantom-hosting-api/config"
	"gitlab.com/jackkdev/phantom-hosting-api/utils"
	"log"
	"net/http"
	"html/template"
)

var mnString string
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("website/*"))
}


func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generateconfigfile", GenerateConfigFile).Methods("POST")
	api.HandleFunc("/generatemasternodestring", GenerateMasternodeString).Methods("POST")
	api.HandleFunc("/addmasternode", AddMasternode).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Frontend Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println("error template")
	}
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

	if mnConfig.TransactionID == "" {
		fmt.Println("A TxID is needed")
		utils.Respond(w, nil, err)
		return
	}

	mnString = config.GenerateNodeDetails(mnConfig)

	utils.Respond(w, mnString, nil)
}

func AddMasternode(w http.ResponseWriter, r *http.Request) {
	config.AddMasternodeToConfigFile("masternode.txt", mnString)
	utils.Respond(w, "Masternode added successfully to configuration file", nil)
}