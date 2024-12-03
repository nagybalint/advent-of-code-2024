package main

import (
	"fmt"
	"log"

	"github.com/nagybalint/advent-of-code-2024/tasks"
)

func main() {
	d := tasks.Day1Task1{}
	answer, err := d.CalculateAnswer()
	if err != nil {
		log.Fatalln("Cannot calculate answer", err)
	}
	fmt.Println(answer)
}
