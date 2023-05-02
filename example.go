package main

import "fmt"

func main() {

	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := append(arr[:3],arr[4:]...)
	fmt.Println(slice)
}