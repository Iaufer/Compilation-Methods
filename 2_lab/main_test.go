package main

import (
	"testing"
)

func TestMC_Correct(t *testing.T) {
	testMap := map[string]string{
		"6 4 7 4 2 3 1":                       "!a+b!",
		"6 7 4 5 2 1":                         "!a*b!",
		"6 4 7 4 2 3 8 7 4 6 4 2 3 8 4 5 2 1": "!(a+b)*(b+a)!",
		"7 6 4 5 6 7 4 5 2 3 1":               "!b*a+a*b!",
		"6 4 7 4 2 3 8 6 4 5 7 6 4 5 2 3 1":   "!(a+b)*a+b*a!",
		"6 4 7 6 4 5 2 3 8 7 7 4 5 6 6 4 7 4 6 4 2 3 3 8 4 5 2 3 8 4 5 2 1": "!(a+b*a)*(b*b+a*(a+b+a))!",
	}

	for key, value := range testMap {
		result := Pars_chain(value)[:len(Pars_chain(value))-1]
		if result != key {
			t.Errorf("Неправильный результат. Ожидалось %s, получили %s", key, result)
		}
	}
}

func TestMC_Incorrect(t *testing.T) {
	testMap := map[string]string{
		"error: !a+*b!":    "!a+*b!",
		"error: a+b*a+b":   "a+b*a+b",
		"error: a!b":       "a!b",
		"error: !a(b+a()!": "!a(b+a()!",
	}

	for key, value := range testMap {
		result := Pars_chain(value)
		if result != "" {
			t.Errorf("Неправильный результат. Ожидалось %s, получили %s", key, result)
		}
	}

}
