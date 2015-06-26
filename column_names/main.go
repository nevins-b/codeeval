package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
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
		value, _ := strconv.Atoi(scanner.Text())
		chars := calculate(value)
		var buf bytes.Buffer
		for _, v := range chars {
			if v == 0 {
				continue
			}
			buf.WriteString(string(rune(64 + v)))
		}
		fmt.Println(buf.String())
	}
}
