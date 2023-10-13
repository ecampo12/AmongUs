package AUInterpreter

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type AUInterpreter struct {
	acc1      int
	acc2      int
	stack     []int8
	currcolor string
	lastwho   int
}

// const (
// 	// commands
// 	SUS = iota
// 	VENTED
// 	SUSSY
// 	ELECTRICAL
// 	WHO
// 	WHERE
// 	// colors
// 	RED
// 	BLUE
// 	PURPLE
// 	GREEN
// 	YELLOW
// 	CYAN
// 	BLACK
// 	WHITE
// 	BROWN
// 	LIME
// 	PINK
// 	ORANGE
// )

var colors = map[string]int{
	"RED":    1,
	"BLUE":   1,
	"PURPLE": 1,
	"GREEN":  1,
	"YELLOW": 1,
	"CYAN":   1,
	"BLACK":  1,
	"WHITE":  1,
	"BROWN":  1,
	"LIME":   1,
	"PINK":   1,
	"ORANGE": 1,
}

func readfile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var code = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// could just do the operations as we read,
		// but loops might be an issue
		// split line into substrings
		if line != "" {
			code = append(code, strings.Split(line, " ")...)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	return code
}

func interpreter(code []string, inputs []string) (string, error) {
	res := ""
	var err error

	// initialize the interpreter
	interpreter := AUInterpreter{0, 0, []int8{}, "", 0}
	// loop through the code
	for i := 0; i < len(code); i++ {
		_, iscolor := colors[code[i]]
		if !iscolor {
			switch code[i] {
			case "SUS":
				color(&interpreter, &res)
			case "VENTED":
				interpreter.acc2 += 10
			case "SUSSY":
				interpreter.acc2 -= 2
			case "ELECTRICAL":
				interpreter.acc1 = 0
			case "WHO", "WHO?":
				interpreter.lastwho = i
				if interpreter.stack[len(interpreter.stack)-1] == int8(interpreter.acc2) {
					for j := i; j < len(code); j++ {
						if code[j] == "WHERE" {
							i = j
							break
						}
					}
				}
			case "WHERE", "WHERE?":
				if interpreter.stack[len(interpreter.stack)-1] != int8(interpreter.acc2) {
					i = interpreter.lastwho
				}
			default:
				err = fmt.Errorf("sus command: %s", code[i])
				return "", err
			}
		} else {
			interpreter.currcolor = code[i]
		}
	}
	return res, err
}

func color(interp *AUInterpreter, res *string) {
	switch interp.currcolor {
	case "RED":
		interp.acc1++
	case "BLUE":
		interp.stack = append(interp.stack, int8(interp.acc1))
	case "PURPLE":
		interp.stack = interp.stack[:len(interp.stack)-1]
	case "GREEN":
		// output the number at the top of stack as a character
		*res += fmt.Sprintf("%c", interp.stack[len(interp.stack)-1])

		// *res += fmt.Sprint(interp.stack[len(interp.stack)-1])
	case "YELLOW":
		// input
	case "CYAN":
		// pop off the stack a random number of times, with highest being interp.acc1
		random := rand.Intn(interp.acc1)
		for i := 0; i < random; i++ {
			interp.stack = interp.stack[:len(interp.stack)-1]
		}
	case "BLACK":
		// output the number at the top of stack as a number
		*res += fmt.Sprintf("%d", interp.stack[len(interp.stack)-1])
	case "WHITE":
		interp.acc1--
	case "BROWN":
		interp.acc1 = int(interp.stack[len(interp.stack)-1])
	case "LIME":
		interp.stack[len(interp.stack)-1] *= 2
	case "PINK":
		interp.acc1 = 0
	case "ORANGE":
		interp.acc1 += 10
	}
}
