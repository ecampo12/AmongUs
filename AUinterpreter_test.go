package AUInterpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	code := readfile("../Testfiles/helloworld")
	output, _ := interpreter(code, nil)
	expected := "Hello World!\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}
