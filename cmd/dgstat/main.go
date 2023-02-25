package main

import (
	"fmt"
	"github.com/rrees/go-cli/v2/internal/dice"
	"sort"
)

func main() {
	statRoll := dice.D(dice.RollRequest{3, 6, 0})

	sort.Ints(statRoll.Results)
	fmt.Printf("Stat: %d (from rolls %v)\n", statRoll.Results[1], statRoll.Results)
}
