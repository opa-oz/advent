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

func extractFromLine(forward, backward string) (int, error) {
	var err error
	left := -1
	right := -1

	for idx := 0; idx < len(forward); idx++ {

		if left != -1 {
			break
		}

		ch := string(forward[idx])
		if strings.Contains(availableValues, ch) {
			left, err = strconv.Atoi(ch)
			if err != nil {
				return 0, err
			}
		}
	}

	for idx := 0; idx < len(backward); idx++ {
		ridx := len(backward) - 1 - idx

		if right != -1 {
			break
		}

		ch := string(backward[ridx])
		if strings.Contains(availableValues, ch) {
			right, err = strconv.Atoi(ch)
			if err != nil {
				return 0, err
			}
		}
	}

	return left*10 + right, nil
}

var replacer = strings.NewReplacer(
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9",
)
var reverseReplacer = strings.NewReplacer(
	"eno", "1",
	"owt", "2",
	"eerht", "3",
	"ruof", "4",
	"evif", "5",
	"xis", "6",
	"neves", "7",
	"thgie", "8",
	"enin", "9",
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func toPartOne(line string) (string, string) {
	forward := replacer.Replace(line)
	backward := reverseReplacer.Replace(Reverse(line))

	return forward, Reverse(backward)
}

func processOneFile(fname string, asPartOne bool) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		forward := line
		backward := line

		if asPartOne {
			forward, backward = toPartOne(line)
		}

		num, err := extractFromLine(forward, backward)
		if err != nil {
			panic(todo.String("Replace with something", "Omg, what to do?!"))
		}

		sum += num
	}

	fmt.Println(fmt.Sprintf("[%s]", fname))
	fmt.Println("\tThe final result is: ", sum)

}

func main() {
	processOneFile(todo.String("Replace with non magic path", "day01/01-test.txt"), false)
	processOneFile(todo.String("meh", "day01/01-input.txt"), false)

	processOneFile(todo.String("meh2", "day01/02-test.txt"), true)
	processOneFile(todo.String("meh3", "day01/02-input.txt"), true)
}
