package tasks

import (
	"strconv"
	"strings"
)

type Day9Task1 struct{}

func (d Day9Task1) CalculateAnswer(input string) (string, error) {

	var files []int
	var memorylayout []int
	for i, rawSize := range strings.TrimRight(input, "\n") {
		size, _ := strconv.Atoi(string(rawSize))
		if i%2 == 0 {
			files = append(files, size)
		}
		memorylayout = append(memorylayout, size)
	}

	leftBlockId := d.takeLeftBlockId(files)
	rightBlockId := d.takeRightBlockId(files)

	checkSum := 0
	blockCounter, memoryId := 0, 0
	for blockCounter < d.totalFileLength(files) {
		if memorylayout[memoryId] == 0 {
			memoryId++
			continue
		}
		var blockIdToAdd int
		if memoryId%2 == 0 {
			blockIdToAdd = <-leftBlockId
		} else {
			blockIdToAdd = <-rightBlockId
		}
		checkSum += blockCounter * blockIdToAdd
		memorylayout[memoryId]--
		blockCounter++
	}

	return strconv.Itoa(checkSum), nil
}

func (d Day9Task1) totalFileLength(files []int) int {
	sum := 0
	for _, fileLength := range files {
		sum += fileLength
	}
	return sum
}

func (d Day9Task1) takeLeftBlockId(files []int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < len(files); i++ {
			for j := 0; j < files[i]; j++ {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func (d Day9Task1) takeRightBlockId(files []int) <-chan int {
	out := make(chan int)
	go func() {
		for i := len(files) - 1; i >= 0; i-- {
			for j := 0; j < files[i]; j++ {
				out <- i
			}
		}
		close(out)
	}()
	return out
}
