// Package constants defines constants used across packages in the project
package constants

type GameOutcome int

const (
	ROCK     = "🪨 Rock"
	PAPER    = "📝 Paper"
	SCISSORS = "✂️ Scissors"
	QUIT     = "Quit"

	WIN  GameOutcome = iota
	LOSS GameOutcome = iota
	TIE  GameOutcome = iota
)

var (
	Moves = []string{ROCK, PAPER, SCISSORS}
)
