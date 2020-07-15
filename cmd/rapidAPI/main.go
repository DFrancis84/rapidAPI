package main

import (
	"fmt"
	"os"

	"github.com/DFrancis84/rapidAPI/internal/covid19"
	"github.com/DFrancis84/rapidAPI/internal/restapi"
	"github.com/DFrancis84/rapidAPI/internal/urbanDictionary"
)

func main() {
	apiKey := os.Getenv("X_RAPIDAPI_KEY")
	urbanDictionary := urbanDictionary.New()
	covid19 := covid19.New()
	api := restapi.New(urbanDictionary, covid19, apiKey)

	fmt.Println("Starting Rapid API Server...")
	api.HandleRequests()
}
