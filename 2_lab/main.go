package main

import (
	"fmt"
	"strconv"
	"strings"
)

// var SyntaxTable = [][]string{
// 	{"?", "!", "+", "*", "(", ")", "a", "b", "$", "A", "B", "T", "M"},
// 	{"0", "S2", "", "", "", "", "", "", "", "1", "", "", ""},
// 	{"1", "", "", "", "", "", "", "", "ex", "", "", "", ""},
// 	{"2", "", "", "", "S3", "", "S4", "S5", "", "", "6", "7", "8"},
// 	{"3", "", "", "", "S3", "", "S4", "S5", "", "", "9", "7", "8"},
// 	{"4", "R6", "R6", "R6", "", "R6", "", "", "", "", "", "", ""},
// 	{"5", "R7", "R7", "R7", "", "R7", "", "", "", "", "", "", ""},
// 	{"6", "S10", "", "", "", "", "", "", "", "", "", "", ""},
// 	{"7", "R2", "S11", "", "", "R2", "", "", "", "", "", "", ""},
// 	{"8", "R4", "R4", "S12", "", "R4", "", "", "", "", "", "", ""},
// 	{"9", "", "", "", "", "S13", "", "", "", "", "", "", ""},
// 	{"10", "", "", "", "", "", "", "", "R1", "", "", "", ""},
// 	{"11", "", "", "", "S3", "", "S4", "S5", "", "", "14", "7", "8"},
// 	{"12", "", "", "", "S3", "", "S4", "S5", "", "", "", "15", "8"},
// 	{"13", "R8", "R8", "R8", "", "R8", "", "", "", "", "", "", ""},
// 	{"14", "R3", "", "", "", "R3", "", "", "", "", "", "", ""},
// 	{"15", "R5", "R5", "", "", "R5", "", "", "", "", "", "", ""},
// }

var SyntaxTable = [][]string{
	{"?", "!", "+", "*", "(", ")", "a", "b", "$", "A", "B", "T", "M"},
	{"0", "S2", "", "", "", "", "", "", "", "1", "", "", ""},
	{"1", "", "", "", "", "", "", "", "ex", "", "", "", ""},
	{"2", "", "", "", "S3", "", "S4", "S5", "", "", "6", "7", "8"},
	{"3", "", "", "", "S3", "", "S4", "S5", "", "", "9", "7", "8"},
	{"4", "R6", "R6", "R6", "", "R6", "", "", "", "", "", "", ""},
	{"5", "R7", "R7", "R7", "", "R7", "", "", "", "", "", "", ""},
	{"6", "S@", "", "", "", "", "", "", "", "", "", "", ""},
	{"7", "R2", "Sq", "", "", "R2", "", "", "", "", "", "", ""},
	{"8", "R4", "R4", "S/", "", "R4", "", "", "", "", "", "", ""},
	{"9", "", "", "", "", "S%", "", "", "", "", "", "", ""},
	{"@", "", "", "", "", "", "", "", "R1", "", "", "", ""},
	{"q", "", "", "", "S3", "", "S4", "S5", "", "", "^", "7", "8"},
	{"/", "", "", "", "S3", "", "S4", "S5", "", "", "", "=", "8"},
	{"%", "R8", "R8", "R8", "", "R8", "", "", "", "", "", "", ""},
	{"^", "R3", "", "", "", "R3", "", "", "", "", "", "", ""},
	{"=", "R5", "R5", "", "", "R5", "", "", "", "", "", "", ""},
}

var Rules = [][]string{
	{"A", "!B!"},
	{"B", "T"},
	{"B", "T+B"},
	{"T", "M"},
	{"T", "M*T"},
	{"M", "a"},
	{"M", "b"},
	{"M", "(B)"},
}

func Pars_chain(str string) {
	arr := make([]string, 3)
	arr_R := make([]string, 0)

	arr[0] = "0"
	arr[1] = str
	i, j := findSymbol(arr[0], string(str[0]))
	arr[2] = SyntaxTable[i][j]

	// fmt.Println(arr)

	for arr[2] != "ex"{//i := 0; i < 13; i++ {
		_, left := "", ""
		if arr[2][0] == 'S' {
			arr = f_S(arr)
		}
		// fmt.Println(arr)

		if len(arr[1]) == 0{
			arr[1] = "$"
		}

		k, j := findSymbol(get_num(arr[0]), string(arr[1][0]))
		// fmt.Println(k, j, SyntaxTable[k][j])
		arr[2] = SyntaxTable[k][j]
		// fmt.Println(arr)
		if arr[2][0] == 'R' {
			arr_R = append(arr_R, string(arr[2][1]))

			arr, left, _ = f_R(arr)
			// fmt.Println(arr, left, right)
		}
		if arr[2][0] == 'G' {
			arr = f_G(left, arr)
			// fmt.Println(arr, left, right)
		}
	}
	fmt.Println(arr_R)
}

func f_G(left string, arr []string) []string {
	Goto := Goto(get_num(arr[0]), left)
	arr[0] += left + Goto
	return arr
}

func f_R(arr []string) ([]string, string, string) {
	index, err := strconv.Atoi(string(arr[2][1]))
	if err != nil {
		fmt.Println("Something went wrong in function f_R!")
		return nil, "", ""
	}
	var left, right string = Rules[index-1][0], Rules[index-1][1]

	arr[0] = arr[0][:len(arr[0])-(len(right)*2)]
	arr[2] = "G"

	return arr, left, right
}


func f_S(arr []string) []string {
	arr[0] += string(arr[1][0]) + arr[2][1:]
	arr[1] = strings.Replace(arr[1], string(arr[1][0]), "", 1)

	return arr
}

func Goto(num, rules string) string {
	i, j := findSymbol(num, rules)
	return SyntaxTable[i][j]
}

func findSymbol(state, symbol string) (int, int) {
	var _i, _j int = 0, 0

	for _, val := range SyntaxTable {
		for j, char := range val {
			if char == symbol {
				_i = j
			}
		}
	}

	for i, _ := range SyntaxTable {
		if SyntaxTable[i][0] == state {
			_j = i
		}
	}
	return _j, _i
}

func get_num(str string) string {
	// tmp := ""
	// for i := 0; i < len(str); i++ {
	// 	if str[i] >= '0' && str[i] <= '9' {
	// 		tmp += string(str[i])
	// 	} else {
	// 		tmp = ""
	// 	}
	// }
	return string(str[len(str)-1])
}


func main() {
	str := "!(a+b*a)*(b*b+a*(a+b+a))!"
	Pars_chain(str)
}
