package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IsDigit(d byte) bool {
	if d >= 48 && d <= 57 {
		return true
	} else {
		return false
	}
}

func Day3Part1() {
	f, err := os.Open("data/day3.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := scanner.Text()
	s += "."
	fmt.Println(s)
	s1 := ""
	for scanner.Scan() {
		s2 := s1
		s1 = s
		s = scanner.Text()
		s += "."
		fmt.Println(s)
		currentDigit := ""
		for i := 0; i < len(s1); i++ {
			if IsDigit(s1[i]) {
				currentDigit += string(s1[i])
			} else {
				if currentDigit != "" {
					found := false
					d, err := strconv.Atoi(currentDigit)
					check(err)
					//left
					j := i - 1
					for j >= 0 && IsDigit(s1[j]) {
						j -= 1
					}
					if j >= 0 && s1[j] != 46 {
						sum += d
						found = true
					}
					if !found {
						//right
						if s1[i] != 46 {
							sum += d
							found = true
						}
					}
					if !found {
						for k := j; k <= j+len(currentDigit)+1; k++ {
							if k >= 0 && k < len(s1) {
								if !IsDigit(s[k]) && s[k] != 46 {
									sum += d
									found = true
									break
								}
							}
						}
					}
					if !found {
						for k := j; k <= j+len(currentDigit)+1; k++ {
							if k >= 0 && k < len(s2) {
								if !IsDigit(s2[k]) && s2[k] != 46 {
									sum += d
									found = true
									break
								}
							}
						}
					}
				}
				currentDigit = ""
			}
		}

	}

	f.Close()
	fmt.Println(sum)
}

func Day3Part2() {
	f, err := os.Open("data/day3.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := scanner.Text()
	s += "."
	s1 := ""
	for scanner.Scan() {
		s2 := s1
		s1 = s
		s = scanner.Text()
		s += "."
		fmt.Println(s1)
		counter := 0
		stars := 0
		for i := 0; i < len(s1); i++ {
			if s1[i] == 42 {
				counter++
				//left
				left := 0
				if i >= 0 && IsDigit(s1[i-1]) {
					left = 1
				}
				right := 0
				if i < len(s1)-1 && IsDigit(s1[i+1]) {
					right = 1
				}
				up := 0
				upl := 0
				upr := 0
				if i >= 0 && i < len(s2) {
					if IsDigit(s2[i]) {
						up = 1
					} else {
						if i-1 >= 0 && IsDigit(s2[i-1]) {
							upl = 1
						}
						if i+1 < len(s2) && IsDigit(s2[i+1]) {
							upr = 1
						}
					}
				}
				down := 0
				downl := 0
				downr := 0
				if i >= 0 && i < len(s) {
					if IsDigit(s[i]) {
						down = 1
					} else {
						if i-1 >= 0 && IsDigit(s[i-1]) {
							downl = 1
						}
						if i+1 < len(s) && IsDigit(s[i+1]) {
							downr = 1
						}
					}
				}
				if left+right+up+down+upl+upr+downl+downr == 2 {
					stars++
					m := 1
					if left == 1 {
						n := ""
						for j := i - 1; j >= 0; j-- {
							if IsDigit(s1[j]) {
								n += string(s1[j])
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if right == 1 {
						n := ""
						for j := i + 1; j < len(s1); j++ {
							if IsDigit(s1[j]) {
								n += string(s1[j])
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if down == 1 {
						n := ""
						for j := i; j < len(s); j++ {
							if IsDigit(s[j]) {
								n += string(s[j])
							} else {
								break
							}
						}
						for j := i - 1; j >= 0; j-- {
							if IsDigit(s[j]) {
								n = string(s[j]) + n
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if downr == 1 {
						n := ""
						for j := i + 1; j < len(s); j++ {
							if IsDigit(s[j]) {
								n += string(s[j])
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if downl == 1 {
						n := ""
						for j := i - 1; j >= 0; j-- {
							if IsDigit(s[j]) {
								n = string(s[j]) + n
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if up == 1 {
						n := ""
						for j := i; j < len(s2); j++ {
							if IsDigit(s2[j]) {
								n += string(s2[j])
							} else {
								break
							}
						}
						for j := i - 1; j >= 0; j-- {
							if IsDigit(s2[j]) {
								n = string(s2[j]) + n
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if upl == 1 {
						n := ""
						for j := i - 1; j >= 0; j-- {
							if IsDigit(s2[j]) {
								n = string(s2[j]) + n
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					if upr == 1 {
						n := ""
						for j := i + 1; j < len(s2); j++ {
							if IsDigit(s2[j]) {
								n += string(s2[j])
							} else {
								break
							}
						}
						t, err := strconv.Atoi(n)
						check(err)
						m *= t
					}
					sum += m
				}
				left = 0
				right = 0
				up = 0
				down = 0
				upl = 0
				upr = 0
				downl = 0
				downr = 0
			}
		}
		if stars < counter {
			fmt.Printf("counter %d, stars %d \n", counter, stars)
		}

	}

	f.Close()
	fmt.Println(sum)
}
