package main

import (
	"fmt"
	"strings"
)

var (
	Alt = map[string]string{
		"A1": "!B!",
		"B1": "T",
		"B2": "T+B",
		"T1": "M",
		"T2": "M*T",
		"M1": "a",
		"M2": "b",
		"M3": "(B)",
	}
	AltIdx = map[string]string{
		"A1": "1",
		"B1": "2",
		"B2": "3",
		"T1": "4",
		"T2": "5",
		"M1": "6",
		"M2": "7",
		"M3": "8",
	}
	Terminals = []string{"a", "b", "(", ")", "!", "+", "*"}
)

func processStateQ(res string, l1 []string, l2 string, i int, n int, state string) ([]string, string, int, string) {
	if l2[0:1] == "a" || l2[0:1] == "b" || l2[0:1] == "(" || l2[0:1] == ")" || l2[0:1] == "!" || l2[0:1] == "+" || l2[0:1] == "*" {
		if l2[0:1] == res[i:i+1] {
			l1 = append(l1, l2[0:1])
			l2 = l2[1:]
			i++
			if i == n {
				if len(l2) == 0 {
					state = "t"
				} else {
					state = "b"
				}
			} else {
				if len(l2) == 0 {
					state = "b"
				}
			}
		} else {
			state = "b"
		}
	} else {
		altL1 := l2[0:1] + "1"
		l1 = append(l1, altL1)
		l2 = strings.Replace(l2, altL1[0:1], Alt[altL1], 1)
	}
	return l1, l2, i, state
}

func processStateB(res string, l1 []string, l2 string, i int, state string) ([]string, string, int, string) {
	if l1[len(l1)-1] == "a" || l1[len(l1)-1] == "b" || l1[len(l1)-1] == "(" || l1[len(l1)-1] == ")" || l1[len(l1)-1] == "!" || l1[len(l1)-1] == "+" || l1[len(l1)-1] == "*" {
		l2 = l1[len(l1)-1] + l2
		l1 = l1[:len(l1)-1]
		i--
	} else {
		altIdxL1 := l1[len(l1)-1]
		newAltIdxL1 := altIdxL1[0:1] + string(altIdxL1[1]+1)
		if Alt[newAltIdxL1] != "" {
			l1 = l1[:len(l1)-1]
			l1 = append(l1, newAltIdxL1)
			l2 = strings.Replace(l2, Alt[altIdxL1], Alt[newAltIdxL1], 1)
			state = "q"
		} else {
			if newAltIdxL1 == "A2" && i == 0 {
				return l1, l2, i, "break"
			}
			l2 = strings.Replace(l2, Alt[altIdxL1], altIdxL1[0:1], 1)
			l1 = l1[:len(l1)-1]
		}
	}
	return l1, l2, i, state
}

func processStateT(l1 []string) string {
	var resRes string
	for _, x := range l1 {
		if AltIdx[x] != "" {
			resRes += " " + AltIdx[x]
		}
	}
	return resRes
}

func startWork(res string) string {
	var l1 []string
	l2 := "A"
	var resRes string
	n := len(res)
	i := 0
	state := "q"
	for {
		if state == "q" {
			l1, l2, i, state = processStateQ(res, l1, l2, i, n, state)
		} else if state == "b" {
			l1, l2, i, state = processStateB(res, l1, l2, i, state)
			if state == "break" {
				break
			}
		} else if state == "t" {
			resRes = processStateT(l1)
			break
		}
	}

	return resRes
	// resultDetails(resRes)
}

func resultDetails(resRes string) {
	if resRes != "" {
		// resRes = resRes[1:]
		fmt.Println("Нисходящий: ", resRes)
		println(len(resRes))
	} else {
		fmt.Println("Введенная строка не может быть выведена с помощью грамматики")
	}
}

func main() {
	i := "!(a+b*a)*(b*b+a*(a+b+a))!"
	i = "!a+*b!"
	res := startWork(i)
	resultDetails(res)
}

//1 2 5 8 3 4 6 2 4 7 4 8 3 4 7 2 4 6
