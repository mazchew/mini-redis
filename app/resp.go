package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Type byte

type RESP struct {
	Type    Type
	Command string
	Input   string
}

// Various RESP kinds
const (
	Integer = ':'
	String  = '+'
	Bulk    = '$'
	Array   = '*'
	Error   = '-'
)

func NewRESP(input []byte) (RESP, error) {
	if len(input) == 0 {
		return RESP{}, fmt.Errorf("Empty Input")
	}
	err := validateEndOfLine(input)
	if err != nil {
		return RESP{}, err
	}

	err = validateType(input[0])
	if err != nil {
		return RESP{}, err
	}

	return handleInput(string(input))

}

func validateEndOfLine(input []byte) error {
	if input[len(input)-2] != '\r' || input[len(input)-1] != '\n' {
		return fmt.Errorf("Invalid end of line")
	}
	return nil
}

func validateType(t byte) error {
	switch t {
	case Integer, String, Bulk, Array, Error:
		return nil
	default:
		return fmt.Errorf("Invalid type")
	}
}

func handleInput(received string) (RESP, error) {
	switch received[0] {
	case Array:
		return handleArray(received)
	default:
		return RESP{}, fmt.Errorf("Invalid type")

	}
}

func handleArray(received string) (RESP, error) {
	split := strings.Split(received, "\r\n")
	if len(split) < 4 {
		return RESP{}, fmt.Errorf("Invalid Array")
	}
	commandLength := strings.TrimPrefix(split[1], "$")
	command := split[2]
	if err := hasSameLength(command, commandLength); err != nil {
		return RESP{}, err
	}
	input := ""
	if len(split) > 4 {
		inputLength := strings.TrimPrefix(split[3], "$")
		input = split[4]
		if err := hasSameLength(input, inputLength); err != nil {
			return RESP{}, err
		}
	}
	return RESP{
		Type:    Array,
		Command: strings.ToUpper(command),
		Input:   input,

	}, nil
}

func hasSameLength(command string, commandLength string) error {
	l, err := strconv.Atoi(commandLength)
	if err != nil {
		return err
	}
	if l != len(command) {
		return fmt.Errorf("Invalid length")
	}
	return nil
}