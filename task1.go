package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const BMIMsg = "Formula used: BMI = weight (kg) / (height (m) * height (m))"

	var name string
	var weight, height, BMI float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]
	fmt.Print("Enter your weight in kg: ")
	fmt.Scanln(&weight)
	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&height)
	BMI = weight / (height * height)
	fmt.Printf("\nWelcome, %s! Let's calculate your BMI.\n", name)
	fmt.Println(BMIMsg)
	fmt.Printf("Your BMI is: %.2f\n", BMI)
	if BMI < 18.5 {
		fmt.Println("Health Status: Underweight")
	} else if BMI >= 18.5 && BMI < 24.9 {
		fmt.Println("Health Status: Normal weight")
	} else if BMI >= 25 && BMI < 29.9 {
		fmt.Println("Health Status: Overweight")
	} else {
		fmt.Println("Health Status: Obese")
	}
}
