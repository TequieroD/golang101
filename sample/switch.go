package main

import "fmt"

func choose(i int) {
    switch i {
    case 1:
        fmt.Println("1")
    case 2, 3:
        fmt.Println("2 or 3")
    case 4:
        fmt.Println("number 4,and keep going.")
        fallthrough
    case 5:
        fmt.Println("5")
    }
}

func anotherChoose(i int) {
    switch  {
    case i > 5:
        fmt.Println(i, "> 5")
    case i < 5:
        fmt.Println(i, "< 5")
    case i == 5:
        fmt.Println(i, "= 5")
    }
}

func main() {

    choose(1)
    choose(3)
    choose(4)

    anotherChoose(6)
}