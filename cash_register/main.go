package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var denoms = map[int]string{
	10000: "ONE HUNDRED",
	5000:  "FIFTY",
	2000:  "TWENTY",
	1000:  "TEN",
	500:   "FIVE",
	200:   "TWO",
	100:   "ONE",
	50:    "HALF DOLLAR",
	25:    "QUARTER",
	10:    "DIME",
	5:     "NICKEL",
	1:     "PENNY",
}

var keys []int

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keys = make([]int, 0, len(denoms))
	for k := range denoms {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		if len(parts) != 2 {
			log.Printf("Found more than two parts!")
			continue
		}
		pp, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			log.Printf("Failed to parse %s", parts[0])
			continue
		}
		c, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Printf("Failed to parse %s", parts[0])
			continue
		}
		diff := int(c*100) - int(pp*100)
		if diff < 0 {
			fmt.Println("ERROR")
			continue
		}

		if diff == 0 {
			fmt.Println("ZERO")
			continue
		}

		var buf bytes.Buffer
		for _, i := range keys {
			j := diff / i
			for k := 0; k < j; k++ {
				buf.WriteString(denoms[i])
				diff = diff - i
				if diff > 0 {
					buf.WriteString(",")
				}
			}
		}
		fmt.Println(buf.String())
	}
}
