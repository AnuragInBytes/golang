package main

import (
	"fmt"
	"math/rand" //small letters and as short as possible
	"math"
)

func main() {
	fmt.Println("Hello, 世界")
	genRandNum()
	fmt.Println(math.Pi) //encapsulation: always access with capital letters if in another package
	fmt.Println(sum(42, 45))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

func genRandNum() {
	fmt.Println("Random number : ", rand.Intn(10))
}

func sum(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // implicitly declaring return values then assign and just return  	DO NOT USE
	x = sum * 4 / 9
	y = sum - x
	return
}