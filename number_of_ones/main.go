package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		bin := strconv.FormatInt(n, 2)
		count := 0
		for _, i := range string(bin) {
			if i == '1' {
				count++
			}
		}
		fmt.Printf("%d\n", count)
	}
}
