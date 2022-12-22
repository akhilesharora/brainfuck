package brainfuck

import (
	"errors"
	"fmt"
	"io"
)

const (
	MinSize = 1
	MaxSize = 30000
)

func ExecOperations(operation string, position int, mem []byte, decodedText *string) int {
	switch operation {
	case ">":
		if position == MaxSize {
			break
		}
		position++
	case "<":
		if position == 0 {
			break
		}
		position--
	case "+":
		mem[position]++
	case "-":
		if mem[position] == 0 {
			mem[position] = 255
			break
		}
		mem[position]--
	case ".":
		character := mem[position]
		*decodedText += fmt.Sprintf("%c", character)
	case ",":
		var ch byte
		_, e := fmt.Scanf("%c", &ch)
		_ = e
		mem[position] = ch
	default:
	}
	return position
}

func Interpret(stream io.Reader) (string, error) {
	if stream == nil {
		return "", errors.New("nil IO reader")
	}
	var final Stack
	var stack Stack
	var operation string
	var decodedText string
	buf := make([]byte, MinSize)
	mem := make([]byte, MaxSize)
	position := 0

	for {
		if len(stack) > 0 {
			operation = stack.Pop()
		} else {
			_, err := io.ReadFull(stream, buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				return "", err
			}
			operation = string(buf)
		}

		switch operation {
		case ">", "<", "+", "-", ".", ",":
			position = ExecOperations(operation, position, mem, &decodedText)
			final.Push(operation)
			break
		case "[":
			final.Push(operation)
			break
		case "]":
			final.Push(operation)
			if mem[position] > 0 {
				interpretClosureCase(&final, &stack)
			}
		default:
		}
	}
	return decodedText, nil
}

func interpretClosureCase(finalStack, stack *Stack) {
	innerLoop := 0
	flag := false
	for {
		operation := finalStack.Pop()
		if operation == "" {
			break
		}
		stack.Push(operation)
		if operation == "]" && flag {
			innerLoop++
		}
		if operation == "[" {
			if innerLoop == 0 {
				break
			} else {
				innerLoop--
			}
		}
		flag = true
	}
}
