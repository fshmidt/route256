package main

import (
	"fmt"
	"unicode"
)

func runeC(a rune) (res bool) {
	if a == 'e' || a == 'u' || a == 'i' || a == 'o' || a == 'a' || a == 'y' || a == 'E' || a == 'U' || a == 'I' || a == 'O' || a == 'A' || a == 'Y' {
		return true
	}
	return false
}

func checkUp(pas *[]rune) {
	var con, vow int
	for _, a := range *pas {
		if unicode.IsUpper(a) {
			return
		} else if unicode.IsLetter(a) {
			if runeC(a) {
				con++
			} else {
				vow++
			}
		}
	}
	if con != 0 {
		*pas = []rune(string(*pas) + "B")
	} else {
		*pas = []rune(string(*pas) + "A")
	}
}

func checkDown(pas *[]rune) {
	var con, vow int
	for _, a := range *pas {
		if unicode.IsLower(a) {
			return
		} else if unicode.IsLetter(a) {
			if runeC(a) {
				con++
			} else {
				vow++
			}
		}
	}
	if con != 0 {
		*pas = []rune(string(*pas) + "c")
	} else {
		*pas = []rune(string(*pas) + "e")
	}
}

func checkVow(pas *[]rune) {
	var vov, con int
	for _, a := range *pas {
		if runeC(a) {
			con++
		} else if unicode.IsLetter(a) {
			vov++
		}
	}
	if con == 0 {
		*pas = []rune(string(*pas) + "y")
	}
	if vov == 0 {
		*pas = []rune(string(*pas) + "b")
	}
}

func checkNum(pas *[]rune) {
	for _, a := range *pas {
		if unicode.IsDigit(a) {
			return
		}
	}
	*pas = []rune(string(*pas) + "8")
}

func main() {
	var nom int
	fmt.Scan(&nom)
	for a := 0; a < nom; a++ {
		var pas string
		fmt.Scan(&pas)
		rPas := []rune(pas)
		checkUp(&rPas)
		checkDown(&rPas)
		checkVow(&rPas)
		checkNum(&rPas)
		fmt.Println(string(rPas))
	}
}
