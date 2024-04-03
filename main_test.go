package main

import (
	"testing"
)

func TestMC_Correct(t *testing.T) {
	testMap := map[string]string{
		"1 3 4 6 2 4 7":                       "!a+b!",
		"1 2 5 6 4 7":                         "!a*b!",
		"1 2 5 8 3 4 6 2 4 7 4 8 3 4 7 2 4 6": "!(a+b)*(b+a)!",
		"1 3 5 7 4 6 2 5 6 4 7":               "!b*a+a*b!",
		"1 3 5 8 3 4 6 2 4 7 4 6 2 5 7 4 6":   "!(a+b)*a+b*a!",
		"1 2 5 8 3 4 6 2 5 7 4 6 4 8 3 5 7 4 7 2 5 6 4 8 3 4 6 3 4 7 2 4 6": "!(a+b*a)*(b*b+a*(a+b+a))!",
	}

	for key, value := range testMap {
		result := startWork(value)[1:]
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
		result := startWork(value)
		if result != "" {
			t.Errorf("Неправильный результат. Ожидалось %s, получили %s", key, result)
		}
	}

}
