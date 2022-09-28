package score

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoring(t *testing.T) {
	testcases := []struct {
		name  string
		roll  []int
		score Score
	}{
		{
			name:  "empty",
			roll:  []int{},
			score: Score{},
		},
		{
			name: "one",
			roll: []int{1},
			score: Score{
				Score:   100,
				Scoring: []int{1},
			},
		},
		{
			name: "straight",
			roll: []int{6, 2, 4, 3, 5, 1},
			score: Score{
				Score:    1500,
				Scoring:  []int{1, 2, 3, 4, 5, 6},
				Straight: true,
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			score := GetScore(tc.roll)
			assert.Equal(t, tc.score.Score, score.Score)
			assert.ElementsMatch(t, tc.score.Scoring, score.Scoring)
			assert.ElementsMatch(t, tc.score.Triples, score.Triples)
			if tc.score.Straight {
				assert.True(t, score.Straight)
			} else {
				assert.False(t, score.Straight)
			}
		})
	}
}
