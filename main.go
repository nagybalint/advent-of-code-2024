package main

import (
	"fmt"
	"io"
	"log"

	"github.com/nagybalint/advent-of-code-2024/tasks"
	"github.com/nagybalint/advent-of-code-2024/utils"
)

func main() {
	log.SetOutput(io.Discard)
	input, err := utils.ReadFileFromRelative("resources/day8.txt")
	if err != nil {
		log.Println("Error reading input")
		panic(err)
	}
	d := tasks.Day8Task2{}
	answer, err := d.CalculateAnswer(input)
	if err != nil {
		log.Fatalln("Cannot calculate answer", err)
	}
	fmt.Println(answer)
}
