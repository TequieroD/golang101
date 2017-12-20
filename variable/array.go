package main

import "fmt"

func main() {
    var x [4]float64
    x[0] = 23
    x[1] = 45
    x[2] = 33
    x[3] = 21

    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))

    // Case 2
    /*
    x := [4]float64{
	    23,
	    45,
	    33,
	    21,
	}
	*/

	// Slice
	slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    slice3 := make([]int, 5)
    copy(slice3, slice2[:len(slice2)-1])
    fmt.Println(slice1, slice2)
    fmt.Println(slice3)
}