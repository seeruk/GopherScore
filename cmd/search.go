package cmd

import (
	"fmt"

	"github.com/SeerUK/GopherScore/modules/wow"
	"github.com/eidolon/gonsole"
)

func SearchCommand() gonsole.Command {
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
		client := wow.NewApiClient(apiKey)
		character, err := client.FindCharacter(region, realm, name)

		if err != nil {
			fmt.Println(err)

			return 1
		}

		resolver := wow.CharacterDataResolver{}

		fmt.Println("Lookup result:")
		fmt.Println("-", character.Name)
		fmt.Println("-", character.Level)
		fmt.Println("-", resolver.ResolveFaction(character.Faction))
		fmt.Println("-", resolver.ResolveRace(character.Race))
		fmt.Println("-", resolver.ResolveClassName(character.Class))

		return 0
	}

	return gonsole.Command{
		Name:        "search",
		Description: "Lookup player information in the command-line.",
		Configure:   configure,
		Execute:     execute,
	}
}
