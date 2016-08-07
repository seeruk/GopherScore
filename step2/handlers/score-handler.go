package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SeerUK/GopherScore/stepx/modules/wow"
	"github.com/gorilla/mux"
)

// ScoreHandler has fields for the dependencies of the score handler.
type ScoreHandler struct {
	calculator wow.ScoreCalculator
	client     *wow.ApiClient
}

// NewScoreHandler create a new instance
func NewScoreHandler(calculator wow.ScoreCalculator, client *wow.ApiClient) *ScoreHandler {
	return &ScoreHandler{
		calculator: calculator,
		client:     client,
	}
}

// Handler handles incoming HTTP requests, providing a score for the given character, on the given
// realm, and in the given region.
func (h *ScoreHandler) Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	region := vars["region"]
	realm := vars["realm"]
	name := vars["name"]

	character, err := h.client.FindCharacter(region, realm, name)

	if err != nil {
		// Assume that if there's an error, it was probably because the character was not found.
		// This of course does not handle all possible error-cases, but it does stop the application
		// from panicking.
		response, err := json.Marshal(Error{
			Code: http.StatusNotFound,
			Message: fmt.Sprintf(
				"Character %s-%s could not be found in the '%s' region.",
				name,
				realm,
				name,
			),
		})

		// This could get to be familiar...
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}

		w.Write(response)
	}

	score := h.calculator.Calculate(*character)

	response, err := json.Marshal(Score{
		Name:    character.Name,
		Realm:   character.Realm,
		Region:  strings.ToUpper(region),
		Level:   character.Level,
		Faction: character.FactionName(),
		Race:    character.RaceName(),
		Class:   character.ClassName(),
		Score:   score,
	})

	// This is duplicated, can you think of a way to improve it?
	if err != nil {
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}

	w.Write(response)
}

// Score represents a character's score.
type Score struct {
	Name    string `json:"name"`
	Realm   string `json:"realm"`
	Region  string `json:"region"`
	Level   int    `json:"level"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Class   string `json:"class"`
	Score   int    `json:"score"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
