package main

import (
	"fmt"
	"github.com/rrees/go-cli/v2/internal/dice"
	"strconv"
)

func main() {
	result := dice.D(dice.RollRequest{1, 3, -3})

	for _, roll := range result.Results {
		fmt.Println("Stat: " + strconv.Itoa(roll))
	}
}
