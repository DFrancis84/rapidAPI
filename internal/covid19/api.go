package covid19

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	covid19types "github.com/DFrancis84/rapidAPI/internal/covid19/types"
)

var (
	apiURL  = "https://covid1910.p.rapidapi.com/data"
	apiHost = "covid1910.p.rapidapi.com"
	count   int
)

type API struct {
	Client *http.Client
}

func New() *API {
	client := &http.Client{}
	return &API{Client: client}
}

func (api *API) GetCountByCountry(apiKey, country, countType, date string) {
	var result covid19types.Count
	fmt.Printf("Retrieving %v count for %v on %v\n", countType, country, date)
	url := fmt.Sprintf("%v/%v/country/%v/date/%v", apiURL, countType, country, date)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	request.Header.Set("x-rapidapi-host", apiHost)
	request.Header.Set("x-rapidapi-key", apiKey)
	res, err := api.Client.Do(request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)

	fmt.Printf("On %v; %v declared %v total %v cases of Covid-19\n", date, strings.ToUpper(country), getCount(result, countType), strings.ToUpper(countType))
}

func getCount(result covid19types.Count, countType string) int {
	for _, x := range result {
		if countType == "confirmed" {
			return x.ConfirmedCount
		}
		if countType == "recovered" {
			return x.RecoveredCount
		}
		if countType == "death" {
			return x.DeathCount
		}
	}
	return 0
}
