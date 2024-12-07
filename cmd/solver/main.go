package main

import (
	"flag"
	"fmt"
	"github.com/jviguy/advent-of-code-2024/internal/solver"
	"log/slog"
)

func main() {
	var day int
	var part int
	var verbose bool
	var help bool

	flag.IntVar(&day, "d", 1, "Provide the day to solve")
	flag.IntVar(&part, "p", 1, "Part of the problem to solve")
	flag.BoolVar(&verbose, "verbose", false, "Verbose")
	flag.BoolVar(&help, "help", false, "Help")

	flag.Parse()
	
	if help {
		flag.PrintDefaults()
		return
	}
	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	slog.Info(fmt.Sprintf("Solving day %d part %d", day, part))
	hander := solver.NewAoCHandler()
	solution, err := hander.SolveDay(day, part)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(fmt.Sprintf("Solution: %d", solution))
}
