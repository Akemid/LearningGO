package main

import (
	"fmt"
	"os"
)


func main(){
	listaArgs := os.Args
	if len(listaArgs) < 2 {
		fmt.Println("Please provide a file name")
		return
	}
	fileName := listaArgs[1]
	fmt.Println(fileName)
	// OPen the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data := make([]byte, 12)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Number of bytes read:", n)
	fmt.Println("Data read:", string(data))
// 	// Close the file
// 	defer file.Close()
// 	// Read the file
// 	data, err := os.ReadFile(fileName)
// 	if err != nil{
// 		fmt.Println("Error:", err)
// 	}
// 	// BYtes to string
// 	text := string(data)
// 	fmt.Println(text)
}