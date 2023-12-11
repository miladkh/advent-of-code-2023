package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day10Part1n2() {
	f, err := os.Open("data/day10.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	sum := 0
	lines := []string{}
	xpos := 0
	ypos := 0
	for scanner.Scan() {
		s := scanner.Text()
		pos := strings.Index(s, "S")
		if pos != -1 {
			xpos = pos
			ypos = len(lines)
		}
		lines = append(lines, s)
	}
	// fmt.Printf("S(x: %d, y: %d)\n", xpos, ypos)
	visited := [][]int{}
	for i := 0; i < len(lines[0]); i++ {
		visited = append(visited, []int{})
		for j := 0; j < len(lines); j++ {
			visited[i] = append(visited[i], -1)
		}
	}

	if isCyclic(xpos, ypos, lines, visited, -1, -1) {
		sum = -1
		for i := 0; i < len(visited); i++ {
			for j := 0; j < len(visited[i]); j++ {
				if visited[i][j] > sum {
					sum = visited[i][j]
				}
			}
		}
		//Part 2
		d := make([][2]int, sum+1)
		for i := 0; i < len(visited); i++ {
			for j := 0; j < len(visited[i]); j++ {
				if visited[i][j] != -1 {
					d[visited[i][j]][0] = i
					d[visited[i][j]][1] = j
				}
			}
		}
		count := 0
		num := len(d)
		for x, v := range visited {
			for y, w := range v {
				if w == -1 {
					j := num - 1
					c := false
					for i := 0; i < num; i++ {
						if (x == d[i][0]) && (y == d[i][1]) {
							// point is a corner
							c = true
							break
						}
						if (d[i][1] > y) != (d[j][1] > y) {
							slope := (x-d[i][0])*(d[j][1]-d[i][1]) - (d[j][0]-d[i][0])*(y-d[i][1])
							if slope == 0 {
								// point is on boundary
								c = true
								break
							}
							if (slope < 0) != (d[j][1] < d[i][1]) {
								c = !c
							}
						}
						j = i
					}
					if c {
						count++
					}
				}
			}
		}

		fmt.Printf("Part 1: %d\n", (sum+1)/2)
		fmt.Printf("Part 2: %d\n", count)
	}

}

func isCyclic(x, y int, g []string, visited [][]int, parentx, parenty int) bool {
	if parentx == -1 {
		visited[x][y] = 0
	} else {
		visited[x][y] = visited[parentx][parenty] + 1
	}
	children := canGo(x, y, g)
	for _, v := range children {
		if visited[v[0]][v[1]] == -1 {
			if isCyclic(v[0], v[1], g, visited, x, y) {
				return true
			}
		} else if v[0] != parentx || v[1] != parenty {
			return true
		}
	}
	return false
}

func canGo(x, y int, g []string) [][2]int {
	result := [][2]int{}
	switch g[y][x] {
	case "S"[0]:
		if x-1 >= 0 && g[y][x-1] == "-"[0] {
			result = append(result, [2]int{x - 1, y})
		}
		if y-1 >= 0 && g[y-1][x] == "|"[0] {
			result = append(result, [2]int{x, y - 1})
		}
		if y+1 < len(g) && g[y+1][x] == "|"[0] {
			result = append(result, [2]int{x, y + 1})
		}
		if x+1 < len(g[0]) && g[y][x+1] == "-"[0] {
			result = append(result, [2]int{x + 1, y})
		}
	case "|"[0]:
		if y-1 >= 0 {
			result = append(result, [2]int{x, y - 1})
		}
		if y+1 < len(g) {
			result = append(result, [2]int{x, y + 1})
		}
	case "-"[0]:
		if x-1 >= 0 {
			result = append(result, [2]int{x - 1, y})
		}
		if x+1 < len(g[0]) {
			result = append(result, [2]int{x + 1, y})
		}
	case "J"[0]:
		if y-1 >= 0 {
			result = append(result, [2]int{x, y - 1})
		}
		if x-1 >= 0 {
			result = append(result, [2]int{x - 1, y})
		}
	case "L"[0]:
		if y-1 >= 0 {
			result = append(result, [2]int{x, y - 1})
		}
		if x+1 < len(g[0]) {
			result = append(result, [2]int{x + 1, y})
		}
	case "7"[0]:
		if x-1 >= 0 {
			result = append(result, [2]int{x - 1, y})
		}
		if y+1 < len(g) {
			result = append(result, [2]int{x, y + 1})
		}
	case "F"[0]:
		if y+1 < len(g) {
			result = append(result, [2]int{x, y + 1})
		}
		if x+1 < len(g[0]) {
			result = append(result, [2]int{x + 1, y})
		}
	}
	return result
}
