package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackkdev/phantom-hosting/config"
	"github.com/jackkdev/phantom-hosting/utils"
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
	r.HandleFunc("/deploy", DeployMasternodeHandler).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generateconfigfile", GenerateConfigFile).Methods("POST")
	api.HandleFunc("/generatemasternodestring", GenerateMasternodeString).Methods("POST")
	api.HandleFunc("/addmasternode", AddMasternode).Methods("POST")
	http.Handle("/", r)

	r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", http.FileServer(http.Dir("www"))))

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

func DeployMasternodeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "deploy.gohtml", nil)
	if err != nil {
		fmt.Println("error template")
	}
}

// API Handlers
func GenerateConfigFile(w http.ResponseWriter, r *http.Request) {
	err := config.GenerateConfigurationFile("masternode.txt")

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, "Configuration file created", nil)
}

func GenerateMasternodeString(w http.ResponseWriter, r *http.Request) {
	var mnConfig config.MasternodeString

	err := json.NewDecoder(r.Body).Decode(&mnConfig)

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	mnString, err = config.GenerateNodeDetails(mnConfig)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, mnString, nil)
}

func AddMasternode(w http.ResponseWriter, r *http.Request) {
	err := config.AddMasternodeToConfigFile("masternode.txt", mnString)

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, "Masternode added successfully to configuration file", nil)
}