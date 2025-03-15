package main

import (
	"fmt"
	"math/rand"
	"mominoes/actin-rock-paper-scizzors/constants"
	"mominoes/actin-rock-paper-scizzors/evaluator"
	"mominoes/actin-rock-paper-scizzors/record"

	"github.com/manifoldco/promptui"
)

var (
	prompt = promptui.Select{
		Label: "Choose an action",
		Items: append(constants.Moves, constants.QUIT),
	}
)

func main() {
	fmt.Println("=================================")
	fmt.Println("Welcome to Rock, Paper, Scissors!")
	fmt.Println("=================================")

	// Create a new record to keep track of wins, losses and ties
	rec := record.NewRecord()

	// Main loop. Each iteration is a single game.
	for i := 1; ; i++ {
		// Prompt the user to make a move
		_, resp, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed: ", err)
			fmt.Println(rec)
			return
		}
		if resp == constants.QUIT {
			break
		}

		// AI makes a random move
		aiResp := constants.Moves[rand.Intn(len(constants.Moves))]

		// Evaluate the game and update the record
		outcome := evaluator.Evaluate(i, resp, aiResp)
		rec.Update(outcome)
	}

	// Summarize the games
	rec.Summarize()
}
