package main

import (
	"bufio" // Bibliothèque permettant de travailler avec les entrées/sorties de buffers
	"fmt"
	"net" // Bibliothèque permettant de travailler avec les réseaux

	// Bibliothèque permettant de travailler avec le système d'exploitation
	"strconv" // Bibliothèque permettant de convertir les chaines de caractères en nombres
	"strings" // Bibliothèque permettant de travailler avec les chaines de caractères
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

// La fonction handleConnection gère les connexions entrantes
func handleConnection(conn net.Conn) {
	fmt.Println("Connexion entrante")
	defer conn.Close() // Ferme la connexion à la fin de la fonction

	// Lit les fichiers texte en tant que paramètres
	reader := bufio.NewReader(conn)

	// Read the incoming data and split it into rows
	data, _ := reader.ReadString('\n')
	dataA := strings.Split(data, ";")[0]
	rowsA := strings.Split(dataA, ",")
	dataB := strings.Split(data, ";")[1]
	rowsB := strings.Split(dataB, ",")
	// Convert the rows into a 2D integer array
	for _, row := range rowsA {
		if row == "" {
			break
		}
		var intRow []int
		for _, value := range strings.Split(row, " ") {
			intValue, _ := strconv.Atoi(value)
			intRow = append(intRow, intValue)
		}
		A = append(A, intRow)
	}
	// Convertit les fichiers texte en matrices entières
	for _, row := range rowsB {
		if row == "" {
			break
		}
		var intRow []int
		for _, value := range strings.Split(row, " ") {
			intValue, _ := strconv.Atoi(value)
			intRow = append(intRow, intValue)
		}
		B = append(B, intRow)
	}
	for i := 0; i < len(A); i++ {
		M = append(M, make([]int, len(B[0])))
	}
	const N = 10
	jobchan := make(chan jobStruct, N)
	resultchan := make(chan [3]int, 2)
	for i := 0; i < 5; i++ {
		go computeCase(resultchan, jobchan)
	}
	go func(PA *[][]int, PB *[][]int) {
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B[0]); j++ {
				job := jobStruct{PA, PB, i, j}
				jobchan <- job
			}
		}
	}(PA, PB)
	for l := 0; l < len(A)*len(B[0]); l++ {
		res := <-resultchan
		M[res[1]][res[2]] = res[0]
	}
	var strMatrix []string
	for _, row := range M {
		var strRow []string
		for _, element := range row {
			strRow = append(strRow, strconv.Itoa(element))
		}
		strMatrix = append(strMatrix, strings.Join(strRow, " "))
	}

	result := strings.Join(strMatrix, "\n")

	fmt.Println(result)
	conn.Write([]byte(result))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		fmt.Println("Connection accepted")
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func computeCase(c chan [3]int, r chan jobStruct) {
	for true {
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
}
