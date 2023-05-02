package main

import "fmt"

func main() {
        arr := []int{10, 20, 90, 70, 60}
        slice := make([]int, 10)
        num := copy(slice, arr)
        fmt.Println(slice)
		fmt.Println(cap(slice))
		fmt.Println(len(arr))
        fmt.Println(num)
}