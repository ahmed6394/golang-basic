// package main

// import "fmt"

// func main(){
// 	arr := [5]int{1,2,3,4,5}
// 	slice := arr[:3]
// 	// fmt.Println(slice, len(slice), cap(slice))
// 	// slice2 := []int{1,2,3}
// 	// fmt.Println(slice2, len(slice2), cap(slice2))
// 	// slice3 := make([]int, 5, 9)
// 	// fmt.Println(slice3, len(slice3), cap(slice3))
// 	slice = append(slice, 6, 7,8)
// 	arr[0] = 10
// 	slice = append(arr[:2], arr[3:]...)
// 	fmt.Println(slice, len(slice), cap(slice))
// 	// fmt.Println(slice2, len(slice2), cap(slice2))
// }

// // func main(){
// 	m := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 		"three": 3,
// 	}
// 	fmt.Println(m)
// 	// delete(m, "two")
// 	value, ok := m["four"]
// 	fmt.Println(value, ok);
// }
