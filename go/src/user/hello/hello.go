package main

import (
	"fmt"

	"user/gramma"
	"user/newmath"
)

func TestVar() {
	var a1 int
	a1 = 1
	var a2, b2 = true, 1
	a3, b3 := 1, false

	fmt.Print(a1)
	fmt.Print(a2)
	fmt.Print(a3)
	fmt.Print(b2)
	fmt.Print(b3)
	fmt.Println()
}

func TestGramma() {
	gramma.BaseGramma()
}

func TestRoutine() {
	gramma.BaseConcurrence()
}

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))

	//wantInput()

	//TestVar()

	//TestGramma()

	TestRoutine()
}

func wantInput() {
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
	var i int
	fmt.Scanf("%d", &i)
}
