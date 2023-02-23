package dice

import (
	"regexp"
)

var NoModiferRegExp = regexp.MustCompile(`^\d+d(\d)+$`)
var DiceStringRegExp = regexp.MustCompile(`^(\d+)d(\d+)[+-](\d+)$`)

func normaliseString(diceString string) string {
	if NoModiferRegExp.MatchString(diceString) {
		return diceString + "+0"
	}

	return diceString
}
