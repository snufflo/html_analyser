package code_analyser

import (
	"fmt"
	"bufio"
	"os"
)

var BRACKET = 1
var OPERATOR = 2
var CONDITIONAL = 3

var brackets = map[string]int {
	"(": BRACKET, ")": BRACKET,
	"{": BRACKET, "}": BRACKET,
	"[": BRACKET, "]": BRACKET,
}

var operators = map[string]bool {
	"+": OPERATOR, "-": OPERATOR, "=": OPERATOR, ":=": OPREATOR, "/": true, "*": true,
	"++": true, "--": true,
}

var conditionals = map[string]bool {
	"==": true, "<": true, ">": true,
}

var logic = map[string]bool {
	"if": true, "else": true, "elif": true,
	"for": true, "switch": true, "case": true,
}


func InitCodeParser() {
	fmt.Println("Enter file name: ")
	reader := bufio.NewReader(os.Stdin)
	file, _ := reader.ReadString('\n')

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	parseCode(data)
}

func parseCode(data byte[]) {
	var type_c int

	for c := range data {
		type_c = isDelimiter(c)
		switch type_c {
		case BRACKET:

		case OPERATOR:

		case CONDITIONAL:
	}
}

func isDelimiter(c rune) (int){
	for b := range brackets {
		if c == b {
			return BRACKET
		}
	}
	for o := range operators {
		if c == b {
			return OPERATOR
		}
	}
	for c := range conditional {
		if c == b {
			return CONDITIONAL
		}
	}
}
