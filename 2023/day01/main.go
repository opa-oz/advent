package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/opa-oz/go-todo/todo"
)

const filename = "day01/01-input.txt"
const availableValues = "1234567890"

func extractFromLine(line string) (int, error) {
	var err error
	left := -1
	right := -1

	for idx := 0; idx < len(line); idx++ {
		ridx := len(line) - 1 - idx

		if left != -1 && right != -1 {
			break
		}

		if left == -1 {
			ch := string(line[idx])
			if strings.Contains(availableValues, ch) {
				left, err = strconv.Atoi(ch)
				if err != nil {
					return 0, err
				}
			}
		}

		if right == -1 {
			ch := string(line[ridx])
			if strings.Contains(availableValues, ch) {
				right, err = strconv.Atoi(ch)
				if err != nil {
					return 0, err
				}
			}
		}
	}

	// fmt.Println("Line: ", line)
	// fmt.Println(fmt.Sprintf("Left: %d, Right: %d", left, right))
	return left*10 + right, nil
}

func processOneFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, err := extractFromLine(line)
		if err != nil {
			panic(todo.String("Replace with something", "Omg, what to do?!"))
		}

		sum += num
	}

	fmt.Println(fmt.Sprintf("[%s]", fname))
	fmt.Println("\tThe final result is: ", sum)

}

func main() {
	processOneFile(todo.String("Replace with non magic path", "day01/01-test.txt"))
	processOneFile(filename)
}
