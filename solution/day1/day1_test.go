package day1

import "testing"

func TestStep(t *testing.T) {
	result1 := fuelRequirements(100756)
	result2 := fuelRequirements(14)
	result3 := fuelRequirements(1969)

	if result1 != 50346 {
		t.Fatal("Result", result1, "doesn't equal expected value of", 50346)
	}

	if result2 != 2 {
		t.Fatal("Result", result2, "doesn't equal expected value of", 2)
	}

	if result3 != 966 {
		t.Fatal("Result", result3, "doesn't equal expected value of", 966)
	}
}