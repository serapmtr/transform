package main

import (
	"fmt"
	"strings"
)

func stringToBinary(s string) string {
	var bin string
	fmt.Println(s)

	for _, c := range s {
		bin += fmt.Sprintf("%b", c)
	}

	return bin
}

func main() {

	var result string
	var input string

	str := make(map[int]string)

	fmt.Scanln(&input)

	bin := stringToBinary(input)
	splitted := strings.Split(bin, "")

	for i, s := range splitted {
		str[i] = s
	}

	for i := 0; i < len(str)-1; i++ {
		if str[i] == "1" {
			result += "0 "

		}
		if str[i] == "0" {
			result += "00 "
		}
		j := 0
		for j < len(str) {
			if str[i] == str[i+j] {
				result += "0"
				j++
			} else {
				break
			}
		}
		i += j - 1

		result += " "
	}

	fmt.Println(result)
}
