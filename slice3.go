package main

import "fmt"

func main() {
        arr := [5]int{10, 20, 90, 70, 60}
        slice := arr[:3]
        fmt.Println(cap(slice))

        slice_2 := make([]int, 5, 20)
        new_slice := append(slice, slice_2...)
        fmt.Println(cap(new_slice))
}