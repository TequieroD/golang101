package main

import (
    "fmt"
    _ "hello"
)

func add(x, y int) int {
    return x + y
}

func change(x, y string) (string, string) {  
    return y, x  
}  

func returnName(x, y string) (greeting, name string) {  
    greeting = x
    name = y
    return  
}  

func main() {
    // Case 1 call add func
    fmt.Println(add(42, 13))

    // Case 2 call change func
    a, b := change("hello", "world")  
    fmt.Println(a, b) 

    // Case 3 call returnName funt
    fmt.Println(returnName("hello", "denny"))
    a, _ = returnName("hello", "denny")
    fmt.Println(a)

}
