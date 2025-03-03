package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand" //small letters and as short as possible
	"runtime"
	"time"
)

var c, python, java bool //variable is not in unintialize, always in zero values(lower bound) OUTSIDE OF FUNCTION

var (
	Tobe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

//conditional statements
func sqrt (x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {

	// switch case
	fmt.Print("Go runs on : ")
	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day after Tomorrow")
	default:
		fmt.Println("Too far away")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// defer
	defer fmt.Println("World")
	fmt.Println("Hello")

	fmt.Println("Ouput of pow function: ")

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("Output of sqrt func: ")
	fmt.Println(sqrt(2), sqrt(-4))

	//loops
	// 1.for loop
	summ := 0
	for i := 0; i < 10; i++ {
		summ += i
	}

	fmt.Println(summ)
	//2. while loop
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println("the sum3 is : ", sum3)

	//3. infinite loop
	sum4 := 1
	for {
		sum4 += sum4
		if(sum4 > 100) {
			break
		}
	}
	fmt.Println("sum4: ", sum4)


	// variables

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

	// fmt print

	fmt.Println("Hello, 世界")

	//functions
	genRandNum()
	fmt.Println(math.Pi) //encapsulation: always access with capital letters if in another package
	fmt.Println(sum(42, 45))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))


	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// pointers
	k , l := 42, 106

	ptr := &k
	fmt.Println("Address fo k: ", ptr)
	fmt.Println("Value of k: ", *ptr)

	*ptr = 36
	fmt.Println("New value of k: ", k)

	ptr = &l
	fmt.Println("valueo of l: ", *ptr)
	*ptr = *ptr / 53
	fmt.Println("New value of l: ",l)

	// struct
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{3, 5})
	v := Vertex{90,89}
	v.X = 69
	fmt.Println(v.X, v.Y)

	v2 := Vertex{1, 2}
	ptrr := &v2
	fmt.Println(ptrr)
	fmt.Println(*ptrr)
	ptrr.X = 1e9
	fmt.Println(v2)
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