package main

import (
	"fmt"

	"github.com/DFrancis84/rapidAPI/config"
	"github.com/DFrancis84/rapidAPI/internal/restapi"
	"github.com/DFrancis84/rapidAPI/internal/urbanDictionary"
)

func main() {
	apiKey := config.GetAPIKey()
	urbanDictionary := urbanDictionary.New()
	api := restapi.New(urbanDictionary, apiKey)

	fmt.Println("Starting Rapid API Server...")
	api.HandleRequests()
}
