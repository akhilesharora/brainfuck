package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akhilesharora/brainfuck/pkg/brainfuck"
)

func main() {
	if len(os.Args) < 1 {
		log.Print("Insufficient arguments to:", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	file, err := brainfuck.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	decodedText, err := brainfuck.Interpret(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decodedText)
}
