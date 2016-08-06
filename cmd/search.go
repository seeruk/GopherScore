package cmd

import (
	"fmt"

	"github.com/SeerUK/GopherScore/modules/wow"
	"github.com/eidolon/gonsole"
)

// SearchCommand has fields for the dependencies of the search command.
type SearchCommand struct {
	calculator wow.ScoreCalculator
}

// NewSearchCommand creates a new instance of SearchCommand.
func NewSearchCommand(calculator wow.ScoreCalculator) *SearchCommand {
	return &SearchCommand{
		calculator,
	}
}

// Command creates an executable search command.
func (c *SearchCommand) Command() gonsole.Command {
	var apiKey string
	var region string
	var realm string
	var name string

	configure := func(d *gonsole.Definition) {
		d.Arg(
			gonsole.StringValue(&apiKey),
			"APIKEY",
			"An API key for the Battle.net API.",
		)

		d.Arg(
			gonsole.StringValue(&region),
			"REGION",
			"A valid server region (one of \"eu\", \"kr\", \"tw\", \"us\").",
		)

		d.Arg(
			gonsole.StringValue(&realm),
			"REALM",
			"A realm name (proper or normalised).",
		)

		d.Arg(
			gonsole.StringValue(&name),
			"NAME",
			"A character name.",
		)
	}

	execute := func() int {
		fmt.Print("Fetching character...")

		client := wow.NewApiClient(apiKey)
		character, err := client.FindCharacter(region, realm, name)

		fmt.Println("Done!")

		if err != nil {
			fmt.Println(err)

			return 1
		}

		fmt.Println(fmt.Sprintf(
			"Score for %s, a level %d %s %s:",
			character.Name,
			character.Level,
			character.FactionName(),
			character.ClassName(),
		))

		fmt.Println("↳ ", c.calculator.Calculate(*character))

		return 0
	}

	return gonsole.Command{
		Name:        "search",
		Description: "Lookup player scores in the command-line.",
		Configure:   configure,
		Execute:     execute,
	}
}
