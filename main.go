package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand" //small letters and as short as possible
)

var c, python, java bool //variable is not in unintialize, always in zero values(lower bound) OUTSIDE OF FUNCTION

var (
	Tobe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	var i int
	fmt.Println(i, c, java, python)
	j := 9
	var z = 10
	var s string = "heo"
	fmt.Println(j, z, s)

	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe);
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var c int
	var f float64
	var d bool
	var st string
	fmt.Printf("%v %v %v %q\n", c, f, d, st)

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