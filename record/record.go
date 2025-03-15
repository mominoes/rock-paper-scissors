// Package record containss a Record struct to keep track of wins, losses and ties
package record

import (
	"fmt"
	"mominoes/actin-rock-paper-scizzors/constants"
)

// Record holds stats about games won, lost etc (by the player).
type Record struct {
	wins   int
	losses int
	ties   int
}

// NewRecord creates a new Record.
func NewRecord() *Record {
	return &Record{
		wins:   0,
		losses: 0,
		ties:   0,
	}
}

func (r *Record) Update(outcome constants.GameOutcome) {
	switch outcome {
	case constants.WIN:
		r.wins++
	case constants.LOSS:
		r.losses++
	case constants.TIE:
		r.ties++
	default:
		fmt.Printf("Invalid outcome update: %v\n", outcome)
	}
}

// Summarize prints a summary of games played in a session
func (r *Record) Summarize() {
	fmt.Println("=================================")
	fmt.Println("Games Summary:")
	fmt.Printf("Played %d, won %d, lost %d, tied %d\n", r.wins+r.losses+r.ties, r.wins, r.losses, r.ties)
	fmt.Println("=================================")
}
