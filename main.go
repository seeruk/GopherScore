package main

import (
	"os"
	"runtime"

	"github.com/SeerUK/GopherScore/cmd"
	"github.com/eidolon/gonsole"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := gonsole.NewApplication("SeerUK/GopherScore", "1.0.0")
	app.AddCommands([]gonsole.Command{
		cmd.SearchCommand(),
	})

	app.Run(os.Args[1:])
}
