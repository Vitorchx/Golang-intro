package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
		fmt.Println("indice: ", i)
		fmt.Println("soma: ", sum)
	}

	fmt.Println(sum)

}
