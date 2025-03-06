package main

import "fmt"

// func main(){
// 	slice := []int{1,2,3,4,5}

// 	for index, value := range slice {
// 		fmt.Printf("Index: %d, Value: %d\n", index, value)
// 	}
// }

func main(){
	result, err := devide(10, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
		fmt.Println(result)
	}
}

func devide(q, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot devide by zero")
	}
	return q/b, nil
}