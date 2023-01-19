package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var A [][]int
var B [][]int

func Read() {
	// Ouvrez le fichier en mode lecture
	fileA, err := os.Open("dataA.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileA.Close()
	fileB, err := os.Open("dataB.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileB.Close()

	// Créez un nouveau scanner pour lire le fichier
	scannerA := bufio.NewScanner(fileA)
	scannerB := bufio.NewScanner(fileB)
	// Pour chaque ligne du fichier...
	for scannerA.Scan() {
		// Obtenez la ligne courante
		line := scannerA.Text()

		// Découpez la ligne en une liste d'éléments
		els := strings.Split(line, " ")

		// Pour chaque élément de la ligne...
		var row []int
		for _, el := range els {
			// Convertir l'élément en entier
			num, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Ajoutez l'élément converti à la ligne
			row = append(row, num)
		}

		// Ajoutez la ligne à la matrice
		A = append(A, row)
	}
	for scannerB.Scan() {
		// Obtenez la ligne courante
		line := scannerB.Text()

		// Découpez la ligne en une liste d'éléments
		els := strings.Split(line, " ")

		// Pour chaque élément de la ligne...
		var row []int
		for _, el := range els {
			// Convertir l'élément en entier
			num, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Ajoutez l'élément converti à la ligne
			row = append(row, num)
		}

		// Ajoutez la ligne à la matrice
		B = append(B, row)
	}

}
func main() {
	Read()
	var wg sync.WaitGroup //  initialise a counter
	fmt.Println(A)
	canal := make(chan [3]int, 2)
	start := time.Now()
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				cumputeCase(canal, A, B, i, j)
			}(i, j)
		}
	}
	for l := 0; l < len(A)*len(A); l++ {
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
