package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var A [][]int
var B [][]int
var M [][]int
var PA *[][]int = &A
var PB *[][]int = &B

type jobStruct struct {
	PA *[][]int
	PB *[][]int
	i  int
	j  int
}

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
	fmt.Print("hello\n")
	for i := 0; i < len(A); i++ {
		M = append(M, make([]int, len(B[0])))
	}
	const N = 10
	jobchan := make(chan jobStruct, N)
	resultchan := make(chan [3]int, 2)
	for i := 0; i < 10; i++ {
		go computeCase(resultchan, jobchan)
	}
	fmt.Print("hello2\n")
	go func(PA *[][]int, PB *[][]int) {
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B[0]); j++ {
				job := jobStruct{PA, PB, i, j}
				jobchan <- job
			}
		}
	}(PA, PB)
	fmt.Print("hello3\n")
	for l := 0; l < len(A)*len(B[0]); l++ {
		res := <-resultchan
		M[res[1]][res[2]] = res[0]
	}
	fmt.Println(M)
}

func computeCase(c chan [3]int, r chan jobStruct) {
	var res [3]int
	var temp = <-r

	var m1 = *(temp.PA)
	var m2 = *(temp.PB)
	res[0] = 0
	res[1] = temp.i
	res[2] = temp.j
	for k := 0; k < len((*(temp.PA))[0]); k++ {
		res[0] += m1[temp.i][k] * m2[k][temp.j]
	}

	c <- res
}
