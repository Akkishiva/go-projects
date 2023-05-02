package main

import "fmt"

func main() {
        arr := [5]int{10, 20, 90, 70, 60}
        slice := arr[:3]
        fmt.Println(cap(slice))
        new_slice := append(slice, 100, 200)
        fmt.Println(cap(new_slice))
}