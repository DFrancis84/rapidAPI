package restapi

import (
	"net/http"

	"github.com/DFrancis84/rapidAPI/internal/covid19"
	"github.com/DFrancis84/rapidAPI/internal/urbanDictionary"
	"github.com/gorilla/mux"
)

type RESTAPI struct {
	UrbanDictionary *urbanDictionary.API
	Covid19         *covid19.API
	APIKey          string
}

func New(urbanDictionary *urbanDictionary.API, covid19 *covid19.API, apiKey string) *RESTAPI {
	return &RESTAPI{
		UrbanDictionary: urbanDictionary,
		Covid19:         covid19,
		APIKey:          apiKey,
	}
}

func (api *RESTAPI) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *RESTAPI) getDefinition(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")
	api.UrbanDictionary.GetDefinition(term, api.APIKey)
}

func (api *RESTAPI) getCountByCountry(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	countType := r.URL.Query().Get("countType")
	date := r.URL.Query().Get("date")
	api.Covid19.GetCountByCountry(api.APIKey, country, countType, date)
}

func (api *RESTAPI) HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.health)
	r.HandleFunc("/covid19", api.getCountByCountry)
	r.HandleFunc("/urbandictionary", api.getDefinition)

	http.ListenAndServe(":8080", r)
}
