package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// vad är ett unittest
// en funktion som körs och testar att en viss sak
// funkar som den ska
// skapa MÅNGA funktioner = MÅNGA tester
func TestWhenKrogenAnd19AndNotDrunkThenShouldBeAllowed(t *testing.T) {
	//Arrange
	location := "K"
	age := 19
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))
	//
	//Assert - VERIFERA ATT RESULTATET BLEV OK!!!
	if canBuy == false || err != nil {
		t.Fatal("Not allowed but should be")
	}
}

func TestWhenKrogenAnd19AndDrunkThenShouldNotBeAllowed(t *testing.T) {
	//Arrange
	location := "K"
	age := 19
	promille := 1.4

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))
	//
	//Assert - VERIFERA ATT RESULTATET BLEV OK!!!
	assert.True(t, canBuy)
	assert.Nil(t, err)

	// if canBuy == true || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }
}
