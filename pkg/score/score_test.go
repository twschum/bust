package score

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoring(t *testing.T) {
	testcases := []struct {
		name  string
		roll  []int
		key   int
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
				Value:   100,
				Scoring: []int{1},
			},
		},
		{
			name: "straight",
			roll: []int{6, 2, 4, 3, 5, 1},
			score: Score{
				Value:    1500,
				Scoring:  []int{1, 2, 3, 4, 5, 6},
				Straight: true,
			},
		},
		{
			name:  "bust",
			roll:  []int{6, 2, 4, 3, 4, 2},
			score: Score{},
		},
		{
			name: "triple",
			roll: []int{2, 2, 2},
			score: Score{
				Value:   200,
				Triples: []int{2, 2, 2},
			},
		},
		{
			name: "triple 5s",
			roll: []int{5, 5, 5},
			score: Score{
				Value:   500,
				Triples: []int{5, 5, 5},
			},
		},
		{
			name: "triple 1s",
			roll: []int{1, 1, 1},
			score: Score{
				Value:   1000,
				Triples: []int{1, 1, 1},
			},
		},
		{
			name: "mixed triple",
			roll: []int{1, 3, 1, 3, 3, 1},
			score: Score{
				Value:   1300,
				Triples: []int{1, 1, 1, 3, 3, 3},
			},
		},
		{
			name: "mixed triple values",
			roll: []int{1, 5, 1, 1, 3},
			score: Score{
				Value:   1050,
				Scoring: []int{5},
				Triples: []int{1, 1, 1},
			},
		},
		{
			name: "triple residule values",
			roll: []int{1, 1, 1, 1, 5, 6},
			score: Score{
				Value:   1150,
				Scoring: []int{1, 5},
				Triples: []int{1, 1, 1},
			},
		},
		{
			name: "triple residule values key",
			roll: []int{1, 5, 3, 3, 3, 6},
			key:  33315,
			score: Score{
				Value:   450,
				Scoring: []int{1, 5},
				Triples: []int{3, 3, 3},
			},
		},
		{
			name: "residule values key",
			roll: []int{1, 5, 5, 4, 3, 6},
			key:  155,
			score: Score{
				Value:   200,
				Scoring: []int{1, 5, 5},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			score := GetScore(tc.roll)
			assert.Equal(t, tc.score.Value, score.Value)
			assert.ElementsMatch(t, tc.score.Scoring, score.Scoring)
			assert.ElementsMatch(t, tc.score.Triples, score.Triples)
			if tc.score.Straight {
				assert.True(t, score.Straight)
			} else {
				assert.False(t, score.Straight)
			}
			if tc.key > 0 {
				assert.Equal(t, tc.key, score.Key())
			}
		})
	}

	t.Run("zero key on empty", func(t *testing.T) {
		s := GetScore([]int{})
		assert.Zero(t, s.Key())
	})

	t.Run("zero key on bust", func(t *testing.T) {
		s := GetScore([]int{2, 4, 4, 3, 6})
		assert.Zero(t, s.Key())
	})
}
