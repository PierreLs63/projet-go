package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Files sent")
	sendFile(conn, "fileA.txt")
	sendFile(conn, "fileB.txt")
	rData := make([]byte, 1024)
	test, _ := conn.Read(rData)
	file, _ := os.Create("test.txt")
	file.Write(rData[:test])
	defer file.Close()
	defer conn.Close()

}
func sendFile(conn net.Conn, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// Read the file line by line and send it to the server
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		_, err = conn.Write(line)
		if err != nil {
			fmt.Println("Error writing to server:", err)
			os.Exit(1)
		}
	}
}
