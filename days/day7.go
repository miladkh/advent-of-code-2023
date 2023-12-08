package days

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	h   string
	bid int
	cat int
}

func Day7Part1() {
	f, err := os.Open("data/day7.txt")
	check(err)
	sum := 0
	hands := []Hand{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		pp := strings.Split(s, " ")
		b, err := strconv.Atoi(pp[1])
		check(err)
		h := pp[0]
		m := map[rune]int{}
		for _, v := range h {
			m[v]++
		}
		cat := 0
		if len(m) == 1 {
			//Five of a kind
			cat = 7
		} else if len(m) == 2 {
			//Four of a kind
			//Full house
			cat = 5
			for _, w := range m {
				if w == 4 {
					cat = 6
					break
				}
			}
		} else if len(m) == 3 {
			//Three of a kind
			//Two pair
			cat = 3
			for _, w := range m {
				if w == 3 {
					cat = 4
					break
				}
			}
		} else if len(m) == 4 {
			//One pair
			cat = 1
		} else {
			//High card
			cat = 0
		}
		h = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "A", "Z"), "K", "Y"), "Q", "X"), "J", "W")
		hands = append(hands, Hand{h: h, bid: b, cat: cat})

	}
	slices.SortFunc(hands, func(a, b Hand) int {
		r := cmp.Compare(a.cat, b.cat)
		if r == 0 {
			r = cmp.Compare(a.h, b.h)
		}
		return r
	})
	for i := 1; i <= len(hands); i++ {
		sum += hands[i-1].bid * i
	}
	fmt.Println(sum)
}

func Day7Part2() {
	f, err := os.Open("data/day7.txt")
	check(err)
	sum := 0
	hands := []Hand{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		pp := strings.Split(s, " ")
		b, err := strconv.Atoi(pp[1])
		check(err)
		h := pp[0]
		m := map[rune]int{}
		for _, v := range h {
			m[v]++
		}
		cat := 0
		if len(m) == 1 {
			//Five of a kind
			cat = 7
		} else if len(m) == 2 {
			//Four of a kind
			//Full house
			if m[74] >= 1 {
				cat = 7
			} else {
				cat = 5
				for _, w := range m {
					if w == 4 {
						cat = 6
						break
					}
				}
			}
		} else if len(m) == 3 {
			//Three of a kind
			//Two pair
			if m[74] >= 2 {
				cat = 6
			} else {
				cat = 3
				for _, w := range m {
					if w == 3 {
						cat = 4
						break
					}
				}
				if m[74] == 1 {
					cat += 2
				}
			}
		} else if len(m) == 4 {
			//One pair
			cat = 1
			if m[74] >= 1 {
				cat = 4
			}
		} else {
			//High card
			cat = 0
			if m[74] == 1 {
				cat = 1
			}
		}
		h = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "A", "Z"), "K", "Y"), "Q", "X"), "J", "1")
		hands = append(hands, Hand{h: h, bid: b, cat: cat})

	}
	slices.SortFunc(hands, func(a, b Hand) int {
		r := cmp.Compare(a.cat, b.cat)
		if r == 0 {
			r = cmp.Compare(a.h, b.h)
		}
		return r
	})
	for i := 1; i <= len(hands); i++ {
		sum += hands[i-1].bid * i
	}
	fmt.Println(sum)
}
