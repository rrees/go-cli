package dice

import (
	"math/rand"
	"time"
)

func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

type RollRequest struct {
	NumberOfDice int
	Sides        int
	Modifier     int
}

type RollResult struct {
	Total   int
	Results []int
}

func D(roll RollRequest) RollResult {
	r := random()
	results := make([]int, roll.NumberOfDice)

	for i := 0; i < roll.NumberOfDice; i++ {
		results[i] = 1 + r.Intn(roll.Sides)
	}

	var total = 0

	for i := 0; i < len(results); i++ {
		total += results[i]
	}

	return RollResult{
		total + roll.Modifier,
		results,
	}
}
