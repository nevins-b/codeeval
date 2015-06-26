package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var coins = []int{1, 3, 5}

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		c, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			panic(err)
		}
		v, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			panic(err)
		}
		var denoms []int
		for _, i := range strings.Split(parts[2], " ") {
			out, err := strconv.ParseInt(i, 10, 0)
			if err != nil {
				panic(err)
			}
			denoms = append(denoms, int(out))
		}

		for i := 0; i <= int(v); i++ {
			fmt.Printf("%d", i)
		}

		fmt.Printf("%d %d %v\n", c, v, denoms)
	}
}
