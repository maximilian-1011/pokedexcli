package pokeapi

import "testing"

func TestAreaLocation(t *testing.T) {
	cases := struct {
		input string
		expected string
	}{
		input: "https://pokeapi.co/api/v2/location-area",
		expected: "canalave-city-area",
	}

	actual, _ := getLocationArea(cases.input)
	if actual.Results[0].Name != cases.expected {
		t.Errorf("Words don't match. expected: '%v' got: '%v'", cases.expected, actual)
	}
}
