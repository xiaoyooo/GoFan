package gramma

import (
	"fmt"
	f "fmt" // f is a alias of fmt
)

func BaseFlow() {
	//r := addR(2)
	//println(r)
	//trycatch()
	reverseOutput()
	f.Println(doublereturn2(1, 2))
	variadicPFunc(1, 2, 3, 4, 5)
}

func condtion() {
	var x int = 1
	if x > 1 {

	}

	// allow define a variable in condition phaser
	if y := 2; y > 1 {

	} else if y < 3 {

	} else {

	}

	i := 1
	switch i {
	case 1:
		println("found 1") // do not need to write 'break'
	case 2:
		println("found 2")
		fallthrough
	case 3:
		println("found 3")
	default:
		println("not found")
	}
}

func gotof() {
	i := 0
Here:
	println(i)
	i++
	goto Here
}

func loop() {
	// for exp1;exp2;exp3 {}
	sum := 0

	for index1 := 1; index1 <= 1000; index1++ {
		sum += index1
	}

	sum = 0
	index2 := 1
	for ; index2 <= 1000; index2++ {
		sum += index2
	}

	sum = 0
	index3 := 1
	for index3 <= 1000 {
		sum += index3
		index3++
	}

	sum = 0
	for index4 := 1; index4 <= 1000; index4++ {
		if index4%2 == 0 {
			continue
		}

		if index4 == 500 {
			break
		}

		sum += index4
	}

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range map1 {
		println(k, v)
	}

	// use _ to abandon var do not want, avoid the compile error
	for _, v := range map1 {
		println(v)
	}

	// there is no while or do while
}

func addR(result int) int {
	wr := result

	defer func() {
		println("result before increment is", wr)
		wr++
	}()

	wr *= 2
	return wr
}

func trycatch() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	println("now peace")
	println("now panic")
	panic(50)
	println("can we go back here?") // no
}

func reverseOutput() {
	for i := 1; i <= 5; i++ {
		defer fmt.Print(" ", i)
	}
}

func doublereturn(a, b int) (int, int) {
	return a + b, a * b
}

// named return param
func doublereturn2(a, b int) (plus int, multiply int) {
	plus, multiply = a+b, a*b

	//return or return plus, multiply
	// we just use return
	return
}

// variadic param
func variadicPFunc(arg ...int) {
	for i, v := range arg {
		fmt.Println(i, v)
	}
}

func transPointer(a *int) {
	*a += 1
}

// func type
type testInt func(int) bool

func isOld(age int) bool {
	return age >= 60
}

func isYoung(age int) bool {
	return age <= 30
}

func testAge(age int, tester testInt) bool {
	return tester(age)
}

func bangbang() {
	testAge(10, isOld)
	testAge(10, isYoung)
}

func main() {
	bangbang()
}
