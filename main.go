package main

import (
	"fmt"
	"sync"
	"time"
)

var a = [][]uint8{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

var b = [][]uint8{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

func main() {
	var wg sync.WaitGroup //  initialise a counter             // one go routine to wait for
	canal := make(chan uint8)
	start := time.Now()
	for i := 0; i < len(a); i++ { //ir
		for j := 0; j < len(a); j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				cumputeCase(canal, a, b, i, j)
			}(i, j)
		}
	}
	for l := 0; l < len(a)*len(a); l++ {
		//fmt.Println(<-canal)
		<-canal
	}
	fmt.Println(time.Since(start))
	wg.Wait() // block until all routines are done executing
}

func cumputeCase(c chan uint8, m1 [][]uint8, m2 [][]uint8, i int, j int) {
	var res uint8 = 0
	for k := 0; k < len(m1[0]); k++ {
		res += m1[i][k] * m2[k][j]
	}
	c <- res
}
