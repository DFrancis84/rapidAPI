package urbanDictionary

import (
	"encoding/json"
	"fmt"
	"net/http"

	dictionarytypes "github.com/DFrancis84/rapidAPI/internal/urbanDictionary/types"
)

var (
	apiURL  = "https://mashape-community-urban-dictionary.p.rapidapi.com/define"
	apiHost = "mashape-community-urban-dictionary.p.rapidapi.com"
)

type API struct {
	Client *http.Client
}

func New() *API {
	client := &http.Client{}
	return &API{Client: client}
}

func (api *API) GetDefinition(term string, apiKey string) {
	result := dictionarytypes.Result{}
	fmt.Printf("Searching Urban Dictionary for '%v'\n", term)
	url := fmt.Sprintf("%v?term=%v", apiURL, term)
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

	json.NewDecoder(res.Body).Decode(&result)
	for i, def := range result.Definitions {
		fmt.Printf("%d: \t%v\n", i+1, def.Definition)
	}
	//fmt.Printf("%+v\n", result)

}
