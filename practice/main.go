package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//type of variable declarion
	// var c,b int = 8, 9
	// var b int8
	// b = 20
	// x := 25.5
	// var name string
	// var age int
	fmt.Println("Enter your data")
	// fmt.Scan(&name, &age)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// fmt.Printf("Name: %s \nage: %d", name, age)
	fmt.Println(input)
}