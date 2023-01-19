package main

import (
	"fmt"
	"sync"
	"time"
)

var a = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

var b = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

func main() {
	var wg sync.WaitGroup //  initialise a counter
	canal := make(chan [3]int, 2)
	start := time.Now()
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				cumputeCase(canal, a, b, i, j)
			}(i, j)
		}
	}
	for l := 0; l < len(a)*len(a); l++ {
		res := <-canal
		fmt.Println(res)
	}
	fmt.Println(time.Since(start))
	wg.Wait() // block until all routines are done executing
}

func cumputeCase(c chan [3]int, m1 [][]int, m2 [][]int, i int, j int) {
	var res [3]int
	res[0] = 0
	for k := 0; k < len(m1[0]); k++ {
		res[0] += m1[i][k] * m2[k][j]
	}
	res[1] = i
	res[2] = j
	c <- res
}
