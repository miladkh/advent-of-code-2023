package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcDiffer(numbers []int) []int {
	newNum := []int{}
	if len(numbers) <= 1 {
		return numbers
	}
	for i := 1; i < len(numbers); i++ {
		newNum = append(newNum, numbers[i]-numbers[i-1])
	}
	return newNum
}

func Day9Part1() {
	f, err := os.Open("data/day9.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		parts := regex.Split(s, -1)
		var numbers = []int{}
		for _, v := range parts {
			n, err := strconv.Atoi(v)
			check(err)
			numbers = append(numbers, n)
		}
		total := [][]int{}
		total = append(total, numbers)
		for {
			numbers = calcDiffer(numbers)
			total = append(total, numbers)
			finish := true
			for _, v := range numbers {
				if v != 0 {
					finish = false
				}
			}
			if finish == true {
				break
			}
		}
		total[len(total)-1] = append(total[len(total)-1], 0)
		for i := len(total) - 2; i >= 0; i-- {
			total[i] = append(total[i], total[i][len(total[i])-1]+total[i+1][len(total[i+1])-1])
		}
		sum += total[0][len(total[0])-1]
	}
	fmt.Println(sum)
}

func Day9Part2() {
	f, err := os.Open("data/day9.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		parts := regex.Split(s, -1)
		var numbers = []int{}
		for _, v := range parts {
			n, err := strconv.Atoi(v)
			check(err)
			numbers = append(numbers, n)
		}
		total := [][]int{}
		total = append(total, numbers)
		for {
			numbers = calcDiffer(numbers)
			total = append(total, numbers)
			finish := true
			for _, v := range numbers {
				if v != 0 {
					finish = false
				}
			}
			if finish == true {
				break
			}
		}
		total[len(total)-1] = append([]int{0}, total[len(total)-1]...)
		for i := len(total) - 2; i >= 0; i-- {
			total[i] = append([]int{total[i][0] - total[i+1][0]}, total[i]...)
		}
		sum += total[0][0]
	}
	fmt.Println(sum)
}
