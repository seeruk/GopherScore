package main

import (
	"os"
	"runtime"

	"github.com/SeerUK/GopherScore/cmd"
	"github.com/SeerUK/GopherScore/modules/wow"
	"github.com/eidolon/gonsole"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	calculator := wow.AggregateScoreCalculator{}
	calculator.AddCalculator(wow.AchievementScoreCalculator{})
	calculator.AddCalculator(wow.ItemsScoreCalculator{})
	calculator.AddCalculator(wow.ProfessionsScoreCalculator{})
	calculator.AddCalculator(wow.ProgressionScoreCalculator{})

	app := gonsole.NewApplication("SeerUK/GopherScore", "1.0.0")
	app.AddCommands([]gonsole.Command{
		cmd.NewSearchCommand(&calculator).Command(),
		cmd.NewServeCommand(&calculator).Command(),
	})

	app.Run(os.Args[1:])
}
