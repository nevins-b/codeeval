package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		for i := len(parts) - 1; i >= 0; i-- {
			buf.WriteString(parts[i])
			if i != 0 {
				buf.WriteString(" ")
			}
		}
		fmt.Println(buf.String())
	}
}
