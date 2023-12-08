package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	f, err := os.Open("data/day2.txt")
	check(err)
	//12 red cubes, 13 green cubes, and 14 blue
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		p1 := strings.Split(s, ":")
		id, err := strconv.Atoi(p1[0][5:])
		check(err)
		games := strings.Split(p1[1], ";")
		notvalid := false
		for g := range games {
			colors := strings.Split(games[g], ",")
			for c := range colors {
				p2 := strings.Split(strings.Trim(colors[c], " "), " ")
				number, err := strconv.Atoi(p2[0])
				check(err)
				switch p2[1] {
				case "blue":
					if number > 14 {
						//not valid
						notvalid = true
					}
				case "red":
					if number > 12 {
						notvalid = true
					}
				case "green":
					if number > 13 {
						notvalid = true
					}
				}
			}
		}
		if notvalid == false {
			sum += id
		}
	}

	f.Close()
	fmt.Println(sum)
}

func Day2Part2() {
	f, err := os.Open("data/day2.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		p1 := strings.Split(s, ":")
		check(err)
		games := strings.Split(p1[1], ";")
		blue := 0
		red := 0
		green := 0
		for g := range games {
			colors := strings.Split(games[g], ",")
			for c := range colors {
				p2 := strings.Split(strings.Trim(colors[c], " "), " ")
				number, err := strconv.Atoi(p2[0])
				check(err)
				switch p2[1] {
				case "blue":
					if number > blue {
						blue = number
					}
				case "red":
					if number > red {
						red = number
					}
				case "green":
					if number > green {
						green = number
					}
				}
			}
		}
		sum += blue * red * green
	}

	f.Close()
	fmt.Println(sum)
}
