package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/SeerUK/GopherScore/stepx/handlers"
	"github.com/SeerUK/GopherScore/stepx/modules/wow"
	"github.com/gorilla/mux"
)

func main() {
	var apiKey string
	var addr string
	var port int

	flag.StringVar(&apiKey, "apikey", "", "An API key for the Battle.net API.")
	flag.StringVar(&addr, "addr", "0.0.0.0", "An address to bind to.")
	flag.IntVar(&port, "port", 8080, "A port to bind to.")
	flag.Parse()

	calculator := wow.AggregateScoreCalculator{}
	calculator.AddCalculator(wow.AchievementScoreCalculator{})
	calculator.AddCalculator(wow.ItemsScoreCalculator{})
	calculator.AddCalculator(wow.ProfessionsScoreCalculator{})
	calculator.AddCalculator(wow.ProgressionScoreCalculator{})

	client := wow.NewApiClient(apiKey)
	handler := handlers.NewScoreHandler(&calculator, client).Handler

	router := mux.NewRouter()
	router.HandleFunc("/score/{region}/{realm}/{name}", handler)

	fmt.Println(fmt.Sprintf("Listening on http://%s:%d/", addr, port))

	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), router)
}
