package main

import (
	"testing"
)

func TestReward(t *testing.T) {
	testData := [][]string{
		[]string{"b:10",		"10 bux"},
		[]string{"part:53",		"P-40 Warhawk part"},
		[]string{"plane:49",		"Hot Air Balloon"},
		[]string{"part:55;b:100",	"Concorde part + 100 bux"},
	}

	for _, data := range testData {
		input := data[0]
		expectedOutput := data[1]
		actualOutput := getHumanReadableReward(input)
		if actualOutput != expectedOutput {
			t.Error("Unexpected reward string: ", actualOutput)
		}
	}
}
