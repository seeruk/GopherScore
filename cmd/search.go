package cmd

import (
	"fmt"

	"github.com/eidolon/gonsole"
)

func SearchCommand() gonsole.Command {
	configure := func(d *gonsole.Definition) {
		// NYI
	}

	execute := func() int {
		fmt.Println("Hello, World!")

		return 0
	}

	return gonsole.Command{
		Name:        "search",
		Description: "Lookup player information in the command-line.",
		Configure:   configure,
		Execute:     execute,
	}
}
