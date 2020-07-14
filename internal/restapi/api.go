package restapi

import (
	"fmt"
	"net/http"

	"github.com/DFrancis84/rapidAPI/config"
	"github.com/DFrancis84/rapidAPI/internal/urbanDictionary"
	"github.com/gorilla/mux"
)

type RESTAPI struct {
	Dictionary *urbanDictionary.API
	Key        *config.APIConfigs
}

func New(urbanDictionary *urbanDictionary.API, apiKey *config.APIConfigs) *RESTAPI {
	return &RESTAPI{
		Dictionary: urbanDictionary,
		Key:        apiKey,
	}
}

func (api *RESTAPI) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *RESTAPI) getDefinition(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")
	api.Dictionary.GetDefinition(term, api.Key)
}

func (api *RESTAPI) HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.health)
	r.HandleFunc("/urbandictionary", api.getDefinition)

	http.ListenAndServe(":8080", r)
	fmt.Println("Starting Rapid API Server...")
}
