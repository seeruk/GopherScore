package cmd

import (
	"net/http"

	"net"

	"fmt"

	"github.com/SeerUK/GopherScore/handlers"
	"github.com/SeerUK/GopherScore/modules/wow"
	"github.com/eidolon/gonsole"
	"github.com/gorilla/mux"
)

// ServeCommand has fields for the dependencies of the serve command.
type ServeCommand struct {
	calculator wow.ScoreCalculator
}

// NewServeCommand creates a new instance of ServeCommand.
func NewServeCommand(calculator wow.ScoreCalculator) *ServeCommand {
	return &ServeCommand{
		calculator: calculator,
	}
}

// Command creates an executable serve command.
func (c *ServeCommand) Command() gonsole.Command {
	var apiKey string

	addr := net.ParseIP("0.0.0.0")
	port := 8080

	configure := func(d *gonsole.Definition) {
		d.Arg(
			gonsole.StringValue(&apiKey),
			"APIKEY",
			"An API key for the Battle.net API.",
		)

		d.Opt(
			gonsole.IPValue(&addr),
			"--addr=ADDR",
			"An address to bind to.",
		)

		d.Opt(
			gonsole.IntValue(&port),
			"--port=PORT",
			"A port to bind to.",
		)
	}

	execute := func() int {
		client := wow.NewApiClient(apiKey)
		handler := handlers.NewScoreHandler(c.calculator, client).Handler

		router := mux.NewRouter()
		router.HandleFunc("/score/{region}/{realm}/{name}", handler)

		host := fmt.Sprintf("%s:%d", addr.String(), port)

		fmt.Println(fmt.Sprintf("Listening on 'http://%s'.", host))

		http.ListenAndServe(host, router)

		return 0
	}

	return gonsole.Command{
		Name:        "serve",
		Description: "Serve player scores via a RESTful web service.",
		Configure:   configure,
		Execute:     execute,
	}
}
