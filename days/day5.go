package days

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var pattern = "\\s+"
var regex = regexp.MustCompile(pattern)

type Day5Line struct {
	source      int
	destination int
	length      int
}

func findDest(lines []Day5Line, source int) int {
	for i := range lines {
		l := lines[i]
		if source >= l.source && source < l.source+l.length {
			return l.destination + source - l.source
		}
	}
	return source
}

func processMap(mps []string) []Day5Line {
	ll := []Day5Line{}
	for i := 0; i < len(mps); i++ {
		ws := regex.Split(mps[i], -1)
		d, err := strconv.Atoi(ws[0])
		s, err := strconv.Atoi(ws[1])
		lg, err := strconv.Atoi(ws[2])
		check(err)
		ll = append(ll, Day5Line{destination: d, source: s, length: lg})
	}
	return ll
}

func Day5Part1() {
	f, err := os.Open("data/day5.txt")
	check(err)
	sum := 0
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	sd := scanner.Text()

	seeds := regex.Split(sd[7:], -1)
	fmt.Println(seeds)
	maps := [][]Day5Line{}
	mps := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			if !IsDigit(s[0]) {
				mps = []string{}
			} else {
				//it's the real mappings
				mps = append(mps, s)
			}
		} else {
			maps = append(maps, processMap(mps))
		}
	}
	sum = math.MaxInt32
	for i := range seeds {
		seed, err := strconv.Atoi(seeds[i])
		check(err)
		r := seed
		for _, v := range maps {
			r = findDest(v, r)
		}
		if r < sum {
			sum = r
		}
	}
	fmt.Println(sum)
}

func mapRange(m [][3]int64, start, end int64) [][]int64 {
	var result [][]int64
	for _, item := range m {
		dst := item[0]
		src := item[1]
		n := item[2]
		if src < end && start < src+n {
			rstart := src
			rend := src + n
			if src < start {
				rstart = start
			}
			if src+n > end {
				rend = end
			}
			result = append(result, []int64{rstart - src + dst, rend - src + dst})
			if start < src {
				result = append(result, mapRange(m, start, src)...)
			}
			if end > src+n {
				result = append(result, mapRange(m, src+n, end)...)
			}
			return result
		}
	}
	result = append(result, []int64{start, end})
	return result
}

func flatten(r [][]int64) [][]int64 {
	slices.SortFunc(r, func(a, b []int64) int { return cmp.Compare(a[0], b[0]) })
	var newRanges [][]int64
	for _, rg := range r {
		if len(newRanges) > 0 && rg[0] <= newRanges[len(newRanges)-1][1] {
			newRanges[len(newRanges)-1][1] = max(newRanges[len(newRanges)-1][1], rg[1])
		} else {
			newRanges = append(newRanges, []int64{rg[0], rg[1]})
		}
	}
	return newRanges
}

func mapRanges(m [][3]int64, s [][]int64) [][]int64 {
	var newRanges [][]int64
	for _, r := range s {
		newRanges = append(newRanges, mapRange(m, r[0], r[1])...)
	}
	return flatten(newRanges)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Day5Part2() {
	f, err := os.Open("data/day5.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	sd := scanner.Text()
	ssss := regex.Split(sd[7:], -1)
	seeds := []int64{}
	for i := 0; i < len(ssss); i++ {
		s1, err := strconv.ParseInt(ssss[i], 0, 64)
		check(err)
		seeds = append(seeds, s1)
	}

	var maps [][][3]int64
	mps := [][3]int64{}
	scanner.Scan()
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			if !IsDigit(s[0]) {
				mps = [][3]int64{}
			} else {
				//it's the real mappings
				ws := regex.Split(s, -1)
				s1, err := strconv.ParseInt(ws[0], 0, 64)
				s2, err := strconv.ParseInt(ws[1], 0, 64)
				s3, err := strconv.ParseInt(ws[2], 0, 64)
				check(err)
				mps = append(mps, [3]int64{s1, s2, s3})
			}
		} else {
			maps = append(maps, mps)
		}
	}

	var s [][]int64
	for i := 0; i < len(seeds); i += 2 {
		s = append(s, []int64{seeds[i], seeds[i] + seeds[i+1]})
	}

	for _, m := range maps {
		s = mapRanges(m, s)
	}

	fmt.Println(s[0][0])
}
