package score

import (
	"sort"

	"golang.org/x/exp/maps"
)

// want a map of roll -> Stats?
// plenty of duplicates exist

type Map map[int]Stats

func (m Map) Insert(s Score) {
	key := s.Key()
	if sc, ok := m[key]; ok {
		sc.Count += 1
		m[key] = sc
	} else {
		m[key] = Stats{Score: s, Count: 1}
	}
}

func (m Map) Values() []Stats {
	return maps.Values(m)
}

type Stats struct {
	Score
	Count int
}

type Score struct {
	Scoring  []int // 1s, 5s, or the straight
	Triples  []int // dice used as a triple
	Value    int   // numeric game score for the roll
	Straight bool  // 1-6
}

func (s *Score) Key() (val int) {
	for _, x := range append(s.Triples, s.Scoring...) {
		val = 10*val + x
	}
	return val
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
		equal := true
		for i, d := range roll {
			if d != i+1 {
				equal = false
				break
			}
		}
		if equal {
			score.Straight = true
			score.Scoring = roll
			score.Value = 1500
			return score
		}
	}
	for i := 0; i < len(roll); i++ {
		d := roll[i]
		// first analyze possible triple from this number
		if len(roll)-i >= 3 && (d == roll[i+1] && d == roll[i+2]) {
			score.Triples = append(score.Triples, roll[i:i+3]...)
			if d == 1 {
				score.Value += 1000
			} else {
				score.Value += 100 * d
			}
			i += 2 // advance over the triple
			continue
		}
		if d == 1 {
			score.Scoring = append(score.Scoring, d)
			score.Value += 100
		}
		if d == 5 {
			score.Scoring = append(score.Scoring, d)
			score.Value += 50
		}
	}
	return score
}
