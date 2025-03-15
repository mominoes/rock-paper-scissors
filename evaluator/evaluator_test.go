package evaluator_test

import (
	"mominoes/actin-rock-paper-scizzors/constants"
	"mominoes/actin-rock-paper-scizzors/evaluator"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name   string
		resp   string
		aiResp string
		want   constants.GameOutcome
	}{
		// Win scenarios
		{
			name:   "Player ROCK beats AI SCISSORS",
			resp:   constants.ROCK,
			aiResp: constants.SCISSORS,
			want:   constants.WIN,
		},
		{
			name:   "Player PAPER beats AI ROCK",
			resp:   constants.PAPER,
			aiResp: constants.ROCK,
			want:   constants.WIN,
		},
		{
			name:   "Player SCISSORS beats AI PAPER",
			resp:   constants.SCISSORS,
			aiResp: constants.PAPER,
			want:   constants.WIN,
		},
		// Loss scenarios
		{
			name:   "Player ROCK loses to AI PAPER",
			resp:   constants.ROCK,
			aiResp: constants.PAPER,
			want:   constants.LOSS,
		},
		{
			name:   "Player PAPER loses to AI SCISSORS",
			resp:   constants.PAPER,
			aiResp: constants.SCISSORS,
			want:   constants.LOSS,
		},
		{
			name:   "Player SCISSORS loses to AI ROCK",
			resp:   constants.SCISSORS,
			aiResp: constants.ROCK,
			want:   constants.LOSS,
		},
		// Tie scenarios
		{
			name:   "Player ROCK ties with AI ROCK",
			resp:   constants.ROCK,
			aiResp: constants.ROCK,
			want:   constants.TIE,
		},
		{
			name:   "Player PAPER ties with AI PAPER",
			resp:   constants.PAPER,
			aiResp: constants.PAPER,
			want:   constants.TIE,
		},
		{
			name:   "Player SCISSORS ties with AI SCISSORS",
			resp:   constants.SCISSORS,
			aiResp: constants.SCISSORS,
			want:   constants.TIE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outcome := evaluator.Evaluate(0, tt.resp, tt.aiResp)
			if outcome != tt.want {
				t.Errorf("Evaluate() got %v, want %v", outcome, tt.want)
			}
		})
	}
}
