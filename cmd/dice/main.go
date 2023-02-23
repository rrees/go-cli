package main

import (
	"fmt"
	"github.com/rrees/go-cli/v2/internal/dice"
	"os"
)

func main() {

	diceString := "1d6"

	arguments := os.Args[1:]

	if len(arguments) == 1 {
		diceString = arguments[0]
	}

	requestedRoll := dice.ParseString(diceString)

	roll := dice.D(requestedRoll)

	fmt.Println(roll)
}
