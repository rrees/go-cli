package main

import (
	"fmt"
	"github.com/rrees/go-cli/v2/internal/dice"
	"strconv"
)

func main() {
	statRoll := dice.D(dice.RollRequest{1, 3, -3})

	fmt.Println("Stat: " + strconv.Itoa(statRoll.Total))
}
