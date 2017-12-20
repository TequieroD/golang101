package main

import "fmt"

var y float64 = 3.3
var (
	x int
	msg string = "golang"
	python bool
)
var z, java = true, "no!"

func main() {
	x = 99
	number := 123

	fmt.Println("x= ", x)
	fmt.Println("y= ", y)
	fmt.Println("z= ", z)
	fmt.Println("msg= ",msg)
	fmt.Println("python= ", python)
	fmt.Println("java=", java)
	fmt.Println("number= ", number)

	fmt.Println(len(msg))
	fmt.Println(msg[0])
	fmt.Printf("%c\n", msg[0])
}
