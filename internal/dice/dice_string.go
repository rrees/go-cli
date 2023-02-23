package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

var NoModiferRegExp = regexp.MustCompile(`^\d+d(\d)+$`)
var DiceStringRegExp = regexp.MustCompile(`^(?P<Size>\d+)d(?P<Sides>\d+)(?P<Modifier>[+-]\d+)$`)

func normaliseString(diceString string) string {
	if NoModiferRegExp.MatchString(diceString) {
		return diceString + "+0"
	}

	return diceString
}

func parseString(diceString string) RollRequest {
	normaliseDiceString := normaliseString(diceString)
	match := DiceStringRegExp.FindStringSubmatch(normaliseDiceString)

	sizeString := match[DiceStringRegExp.SubexpIndex("Size")]

	poolSize, err := strconv.Atoi(sizeString)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse %s, %s was not an integer", diceString, sizeString))
	}

	sidesString := match[DiceStringRegExp.SubexpIndex("Sides")]
	sides, err := strconv.Atoi(sidesString)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse %s, %s was not an integer", diceString, sidesString))
	}

	modifierString := match[DiceStringRegExp.SubexpIndex("Modifier")]
	modifier, err := strconv.Atoi(modifierString)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse %s, %s was not an integer", diceString, modifierString))
	}

	return RollRequest{
		poolSize,
		sides,
		modifier,
	}
}