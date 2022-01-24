package gocolor_test

import (
	coloring "github.com/godzillaframework/gocolor"
	"testing"
)

func testPrint() {
	coloring.Print("Hello").In("green")
}

func tests(t *testing.T) {
	str := "WAIT WHY THIS CODE LOOKS GOOD??"

	expected := "\033[32m EXPECTED \033[0m\n"

	if str != expected {
		t.Error("Expected", expected)
	}
}
