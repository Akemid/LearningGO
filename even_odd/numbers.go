package main

import "fmt"


type numbers []int

func createNumbers(i int) numbers{
	numbers  := make([]int,i)

	for i := range numbers {
		numbers[i] = i+1
	}
	return numbers
}

func (n numbers) print() {
	for _, number := range n {
		if number % 2 == 0 {
			fmt.Printf("%v Es un numero par \n",number)
		}else{
			fmt.Printf("%v Es un numero impar \n", number)
		}
		
	}
}