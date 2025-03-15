// Package evaluator contains utilities to evaluate the output of a single rock-paper-scizzors game
package evaluator

import (
	"fmt"
	"mominoes/actin-rock-paper-scizzors/constants"
)

// Evaluate compares the response of the player and the AI
func Evaluate(i int, resp, aiResp string) constants.GameOutcome {
	fmt.Printf("Game %d: You chose %s while AI chose %s\n", i, resp, aiResp)

	if resp == aiResp {
		fmt.Println("It's a tie!")
		fmt.Println("-----------------")
		return constants.TIE

	} else if (resp == constants.ROCK && aiResp == constants.SCISSORS) ||
		(resp == constants.PAPER && aiResp == constants.ROCK) ||
		(resp == constants.SCISSORS && aiResp == constants.PAPER) {
		fmt.Println("You win!")
		fmt.Println("-----------------")
		return constants.WIN

	} else {
		fmt.Println("You lose!")
		fmt.Println("-----------------")
		return constants.LOSS
	}
}
