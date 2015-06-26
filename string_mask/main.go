package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculate(value int) []int {
	c := value
	chars := []int{0, 0, 0}
	for {
		if c-26 <= 0 {
			chars[2] = c
			break
		} else {
			c = c - 26
			chars[1]++
			if chars[1] > 26 {
				chars[0]++
				chars[1] = 1
			}
		}
	}
	return chars
}

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		var buf bytes.Buffer
		s := []rune(parts[0])
		mask := []rune(parts[1])
		for i, c := range s {
			value, _ := strconv.Atoi(string(mask[i]))
			if value == 1 {
				buf.WriteString(strings.ToUpper(string(c)))
				continue
			}
			buf.WriteRune(c)
		}
		fmt.Println(buf.String())
	}
}
