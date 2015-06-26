package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var faceValues = map[string]int64{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var combination = map[string]int64{
	"high":     0,
	"pair":     1,
	"2pair":    2,
	"three":    3,
	"straight": 4,
	"flush":    5,
	"full":     6,
	"four":     8,
	"sFlush":   9,
	"rFlush":   10,
}

type card struct {
	value    int64
	facecard bool
	face     string
	suite    string
}

type hand struct {
	cards       []*card
	combination int
	high        int64
}

func (h *hand) minCard() *card {
	v := h.cards[0]
	for _, c := range h.cards {
		if c.value < v.value {
			v = c
		}
	}
	return v
}

func (h *hand) score() int64 {
	matchingSuites := make(map[string]int64)
	matchingValues := make(map[int64]int64)
	for _, card := range h.cards {
		if card.value > h.high {
			h.high = card.value
		}
		v, ok := matchingValues[card.value]
		if !ok {
			matchingValues[card.value] = 1
		} else {
			matchingValues[card.value] = v + 1
		}
		s, ok := matchingSuites[card.suite]
		if !ok {
			matchingSuites[card.suite] = 1
		} else {
			matchingSuites[card.suite] = s + 1
		}
	}

	var currentScore int64
	currentScore = 0
	//fmt.Printf("Suites: %v Values: %v\n", matchingSuites, matchingValues)
	for _, matches := range matchingSuites {
		if matches == 5 {
			var cValues []int64
			faceCount := 0
			for _, c := range h.cards {
				cValues = append(cValues, c.value)
				if c.facecard {
					faceCount++
				}
			}
			if faceCount == 5 {
				// Five face cards, Royal flush!
				currentScore = combination["rFlush"]
				break
			}
			min := h.minCard()
			for i := 0; i < len(h.cards); i++ {
				if i == 4 {
					// Five consecutive cards, Straight Flush!
					currentScore = combination["sFlush"]
					break
				}
				found := false
				for _, c := range h.cards {
					// next consecutive card
					if min.value+1 == c.value {
						min = c
						found = true
					}
				}
				if !found {
					break
				}
			}
			// Five cards of the same suite, Flush!
			currentScore = combination["flush"]
		}
	}

	count := make(map[int64]int64)
	for _, matches := range matchingValues {
		c, ok := count[matches]
		if !ok {
			count[matches] = 1
		} else {
			count[matches] = c + 1
		}
	}
	_, ok := count[4]
	if ok {
		// Four of a kind!
		if currentScore < combination["four"] {
			currentScore = combination["four"]
		}
	}

	_, ok = count[3]
	if ok {
		_, ok = count[2]
		if ok {
			// Three of a kind and a pair, Full House!
			if currentScore < combination["full"] {
				currentScore = combination["full"]
			}
		}
		// Three of a kind!
		if currentScore < combination["three"] {
			currentScore = combination["three"]
		}
	}
	t, ok := count[2]
	if ok {
		if t == 2 && currentScore < combination["2pair"] {
			currentScore = combination["2pair"]
		} else if currentScore < combination["pair"] {
			currentScore = combination["pair"]
		}
	}
	return currentScore
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
		cards := strings.Split(scanner.Text(), " ")
		left := &hand{
			high: 0,
		}
		right := &hand{
			high: 0,
		}
		for i, c := range cards {
			def := strings.Split(c, "")
			v := &card{
				suite: def[1],
			}
			value, err := strconv.ParseInt(def[0], 10, 64)
			if err != nil {
				if def[0] != "T" && def[0] != "A" {
					v.facecard = true
				}
				v.face = def[0]
				v.value, _ = faceValues[def[0]]
			} else {
				v.value = value
				v.facecard = false
			}
			if i < 5 {
				left.cards = append(left.cards, v)
			} else {
				right.cards = append(right.cards, v)
			}
		}
		l := left.score()
		r := right.score()
		if l > r {
			fmt.Println("left")
		} else if r > l {
			fmt.Println("right")
		} else {
			if left.high > right.high {
				fmt.Println("left")
			} else if right.high > left.high {
				fmt.Println("right")
			} else {
				fmt.Println("none")
			}
		}
	}
}
