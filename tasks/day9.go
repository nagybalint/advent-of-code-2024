package tasks

import (
	"container/list"
	"fmt"
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
	memoryAddress, itemId := 0, 0
	for memoryAddress < d.totalFileLength(files) {
		if memorylayout[itemId] == 0 {
			itemId++
			continue
		}
		var blockIdToAdd int
		if itemId%2 == 0 {
			blockIdToAdd = <-leftBlockId
		} else {
			blockIdToAdd = <-rightBlockId
		}
		checkSum += memoryAddress * blockIdToAdd
		memorylayout[itemId]--
		memoryAddress++
	}

	return strconv.Itoa(checkSum), nil
}

type Day9Task2 struct{}
type memoryItem interface {
	Size() int
}

type file struct {
	BlockId int
	size    int
}

func (f file) Size() int {
	return f.size
}

type space struct {
	size int
}

func (s space) Size() int {
	return s.size
}

func (d Day9Task2) CalculateAnswer(input string) (string, error) {
	memory := list.New()
	blockId := 0
	for i, rawSize := range strings.TrimRight(input, "\n") {
		size, _ := strconv.Atoi(string(rawSize))
		if i%2 == 0 {
			memory.PushBack(file{BlockId: blockId, size: size})
			blockId++
		} else {
			memory.PushBack(space{size: size})
		}
	}

	for itemToDefragment := memory.Back(); itemToDefragment != memory.Front(); itemToDefragment = itemToDefragment.Prev() {
		if _, isSpace := itemToDefragment.Value.(space); isSpace {
			continue
		}
		fileToDefragment := file(itemToDefragment.Value.(file))
	defragmentItem:
		for defragmentLocation := memory.Front(); defragmentLocation != itemToDefragment; defragmentLocation = defragmentLocation.Next() {
			switch defragmentLocation.Value.(type) {
			case file:
				continue
			case space:
				spaceToDefragmentInto := space(defragmentLocation.Value.(space))
				if spaceToDefragmentInto.Size() < fileToDefragment.Size() {
					continue
				}
				// We have found a space to defragment into
				// Move file forward, adjust spaces
				memory.InsertBefore(space{size: 0}, defragmentLocation)
				memory.InsertBefore(fileToDefragment, defragmentLocation)
				memory.InsertBefore(space{size: spaceToDefragmentInto.Size() - fileToDefragment.Size()}, defragmentLocation)
				memory.Remove(defragmentLocation)

				// Remove file from back, reconsolidate spaces, set itemToDefragment to the empty file that was added instead
				// so that the next iteration will continue from the item before that
				memory.InsertAfter(file{BlockId: 0, size: 0}, itemToDefragment)
				memory.InsertAfter(space{size: fileToDefragment.size}, itemToDefragment)
				memory.InsertAfter(file{BlockId: 0, size: 0}, itemToDefragment)
				replacementFile := itemToDefragment.Next()
				memory.Remove(itemToDefragment)
				itemToDefragment = replacementFile

				// Finished defragmentation of element, element is in its final location
				// Break out of the inner for loop to continue with defragmenting the next element
				break defragmentItem
			default:
				continue
			}
		}

	}

	checkSum := 0
	memoryAddress := 0
	for e := memory.Front(); e != nil; e = e.Next() {
		if f, isFile := e.Value.(file); isFile {
			checkSum += f.size*memoryAddress*f.BlockId + f.BlockId*f.size*(f.size-1)/2
		}
		memoryAddress += memoryItem(e.Value.(memoryItem)).Size()
	}

	return strconv.Itoa(checkSum), nil
}

func memoryAsString(memory list.List) string {
	var sb strings.Builder
	for e := memory.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case file:
			f := file(e.Value.(file))
			for i := 0; i < f.Size(); i++ {
				sb.WriteString(fmt.Sprintf("%d", f.BlockId))
			}
		case space:
			s := space(e.Value.(space))
			for i := 0; i < s.Size(); i++ {
				sb.WriteString(".")
			}
		default:
			continue
		}
	}
	return sb.String()
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
