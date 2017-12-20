package main

import "fmt"

func main() {
    i := 0
HERE:
    fmt.Print(i, " ")
    i++
    if i < 10 {
        goto HERE
    }
}