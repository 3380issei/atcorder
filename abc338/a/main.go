package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	if !IsCapital(rune(s[0])) {
		fmt.Println("No")
		return
	}

	for i := 1; i < len(s); i++ {
		if IsCapital(rune(s[i])) {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func IsCapital(r rune) bool {

	capital := []rune{
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',
	}

	for _, c := range capital {
		if c == r {
			return true
		}
	}

	return false
}
