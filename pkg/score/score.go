package score

import "sort"

// want a map of roll -> Stats?
// plenty of duplicates exist

type Score struct {
	Scoring  []int // 1s, 5s, or the straight
	Triples  []int // dice used as a triple
	Score    int   // numeric game score for the roll
	Straight bool  // 1-6
}

// GetScore leaves only the scoring dice in the slice
// assumes you always "promote" a triple 1s or 5s
// scoring are lone 1s and 5s (or maybe the straight?)
func GetScore(roll []int) Score {
	score := Score{}
	// checking the triple needs the roll to be sorted
	sort.Slice(roll, func(i, j int) bool {
		return roll[i] < roll[j]
	})
	// check for a straight
	if len(roll) == 6 {
		//roll == []int{1,2,3,4,5,6} {
		score.Straight = true
		score.Scoring = roll
		score.Score = 1500
		return score
	}
	for i := 0; i < len(roll); i++ {
		d := roll[i]
		// first analyze possible triple from this number
		if len(roll)-i >= 3 && (d == roll[i+1] && d == roll[i+2]) {
			score.Triples = append(score.Triples, roll[i:i+3]...)
			if d == 1 {
				score.Score += 1000
			} else {
				score.Score += 100 * d
			}
			i += 2 // advance over the triple
			continue
		}
		if d == 1 {
			score.Scoring = append(score.Scoring, d)
			score.Score += 100
		}
		if d == 5 {
			score.Scoring = append(score.Scoring, d)
			score.Score += 50
		}
	}
	return score
}
