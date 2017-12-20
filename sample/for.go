package main

import "fmt"

func main() {
	// Case 1
	// 1 + 2 + 3 + .. + 10
    sum := 0
    for i := 0; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    // Case 2
    // 1 + 2 + 4 + ...
    sum = 1
    for ; sum <= 100; {
        sum += sum
    }
    fmt.Println(sum)

    // Same as Case 2
    sum = 1
    for sum <= 100 {
        sum += sum
    }
    fmt.Println(sum)

    // Infinite Loop
    for {
        }	
}