package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
		n, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			panic(err)
		}

		count := 0
		v := int(n)
		for _, i := range coins {
			j := v / i
			for k := 0; k < j; k++ {
				count++
				v = v - i
			}
		}
		fmt.Printf("%d\n", count)
	}
}
