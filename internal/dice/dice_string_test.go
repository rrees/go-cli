package dice

import (
	"fmt"
	"testing"
)

func Test_diceStringNormalisationForMissingModifiers(t *testing.T) {
	inputString := "1d6"
	result := normaliseString(inputString)

	if result != "1d6+0" {
		t.Error("Dice string ", inputString, " not normalised, got ", result)
	}
}

func Test_diceStringNormalisationForProvidedModifiers(t *testing.T) {
	inputString := "1d6+2"
	result := normaliseString(inputString)

	if result != "1d6+2" {
		t.Error("Dice string ", inputString, " not normalised, got ", result)
	}
}

func Test_diceStringCreatesRollRequest(t *testing.T) {
	for _, testData := range []struct{
		Input string
		ExpectedPool int
		ExpectedSides int
		ExpectedModifier int
	}{
		{"11d20-4", 11, 20, -4,},
		{"1d6", 1, 6, 0},
		{"2d4+1", 2, 4, 1},
	} {
		result := parseString(testData.Input)

		if result.NumberOfDice != testData.ExpectedPool {
			t.Error(fmt.Sprintf("Wrong number of dice read: expected %d and got %d ", testData.ExpectedPool, result.NumberOfDice))
		}

		if result.Sides != testData.ExpectedSides {
			t.Error(fmt.Sprintf("Wrong number of sides requested, expected %d and got %d", testData.ExpectedSides, result.Sides))
		}

		if result.Modifier != testData.ExpectedModifier {
			t.Error(fmt.Sprintf("Wrong modifer requested, expected %d and got %d", testData.ExpectedModifier, result.Modifier))
		}
	}

}