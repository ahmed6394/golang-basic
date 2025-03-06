package main

import (
	// "bufio"
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
	// fmt.Println("Enter your data")
	// fmt.Scan(&name, &age)
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// fmt.Printf("Name: %s \nage: %d", name, age)
	// fmt.Println(input)

	const BD = "BD"
	var name, country, message string
	var age int
	fmt.Print("Enter your name: \n")
	fmt.Scanln(&name)
	fmt.Print("Enter your country: \n")
	fmt.Scanln(&country)
	fmt.Print("Enter your age: \n")
	fmt.Scanln(&age)
	if country == BD && age >= 18 {
		message = fmt.Sprintf("Hello %s! \nYour age is: %d \nVoting status: Eligible", name, age)
	} else {
		message = fmt.Sprintf("Hello %s! \nYour age is: %d \nVoting status: Not Eligible", name, age)
	}
	file, _ := os.Create("output.txt")
	fmt.Fprintf(file, message)
	file.Close()
	fmt.Print("Your voting status has been saved in output.txt file")
}