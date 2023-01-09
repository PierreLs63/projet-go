package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup
const N int = 3

var a = [N][N]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}
var b = [N][N]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

func run(i int, j int) {
	defer wg.Done()
	for i := 0; i < N; i++ {

	}
}

func main() {
	debut := time.Now()
	for i := 0; i < N; i++ {
		for j := 0; j < N; i++ {
			wg.Add(1)
			go run(i, j)
		}
	}

	wg.Wait()
	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}
