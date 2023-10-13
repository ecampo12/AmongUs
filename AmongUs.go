package AUInterpreter

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Print("Usage: go run Brainfuck.go <file> <input values>")
	os.Exit(1)
}

func AmongUs() {
	if len(os.Args) < 2 {
		usage()
	}

	filename := os.Args[1]

	code := readfile(filename)

	var inputs []string
	if len(os.Args) > 2 {
		inputs = os.Args[2:]
		fmt.Println(inputs)
	}

	res, err := interpreter(code, inputs)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(res)

}
