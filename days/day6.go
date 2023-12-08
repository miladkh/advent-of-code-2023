package days

import (
	"fmt"
	"math"
)

type PairInt struct {
	a int
	b float64
}

func Day6Part1() {
	// Time:        44     89     96     91
	// Distance:   277   1136   1890   1768
	sum := 1
	// input := []PairInt{PairInt{a: 44, b: 277}, PairInt{a: 89, b: 1136}, PairInt{a: 96, b: 1890}, PairInt{a: 91, b: 1768}}
	input := []PairInt{PairInt{a: 44899691, b: 277113618901768}}
	for i := 0; i < len(input); i++ {
		t := input[i].a
		d := input[i].b
		counter := 0
		for i := 1; i < t-1; i++ {
			total := i + int(math.Ceil((d / float64(i))))
			if total <= t {
				counter++
			}
		}
		fmt.Println(counter)
		sum *= counter
	}
	fmt.Println(sum)

}
