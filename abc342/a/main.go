package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	majority := s[0]
	if s[1] == s[2] {
		majority = s[1]
	}

	for i, c := range s {
		if c != rune(majority) {
			fmt.Println(i + 1)
			return
		}
	}
}
