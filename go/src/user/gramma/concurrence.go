package gramma

import (
	"fmt"
	"runtime"
	//"bufio"
	//"os"
	//"reflect"
	"strconv"
)

func BaseConcurrence() {
	//goRun1_0()
	//goRun1_1()
	//goRun2()

	goRun2_1()

	gorun3_0()

	gorun3_1()

	//wantInput()
}

func wantInput() {
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
	fmt.Println("now we wait")
	var i int
	fmt.Scanf("%d", &i)
}

// go routine
// without call runtime.GOMAXPROCS, the go routine only run in one thread,
// the result will be very different with the multi thread.
// call runtime.GOMAXPROCS will change the size of the schedulers thread pool
var inc int = 1

func RePrint1(s string) {
	fmt.Print(s + ", ")
	inc++
}

func trigsl() int {
	return runtime.GOMAXPROCS(1)
}

func trigml() int {
	return runtime.GOMAXPROCS(2)
}

func goRun1_1() {
	fmt.Println("in multi threads ", runtime.GOMAXPROCS(2))
	goRun1()
}

// but even in one thread, go routine also can be scheduled to run concurrently
func goRun1_0() {
	fmt.Println("only in one thread")
	goRun1()
}

func goRun1() {
	inc = 1

	fmt.Println("that's begin #1")
	const size = 100
	ss := make([]int, size)

	for i := 1; i <= size; i++ {
		ss[i-1] = i
	}

	// it will almost do not display 1-100 in order
	for _, s := range ss {
		go RePrint1(strconv.Itoa(s))
	}

	wantInput()
	fmt.Println("that's end #1 ", inc)
}

func RePrint2(s string, c chan string) {
	c <- s
	inc++
}

func goRun2() {
	fmt.Println("that's begin #2")
	const size = 50
	ss := make([]int, size)
	inc = 1

	for i := 1; i <= size; i++ {
		ss[i-1] = i
	}

	// default chan is block
	// c := make(chan string)
	// , but we can modify the window size of chan(buffered chan).
	c := make(chan string, 5)
	for _, s := range ss {
		go RePrint2(strconv.Itoa(s), c)

		l := <-c
		fmt.Print(l + ". ")
	}

	wantInput()
	fmt.Println("that's end #2 ", inc)
}

func fibonacci(n int, c chan int) int {
	defer func() {
		close(c)
	}()

	x, y := 1, 1
	for i := 1; i <= n; i++ {
		x, y = y, x+y
		c <- y
	}

	return y
}

func goRun2_1() {
	fmt.Println("that's begin #2.1")

	const fibsize = 10

	c := make(chan int, fibsize)
	go fibonacci(fibsize, c)

	for s := range c {
		fmt.Print(s, ",")
	}

	wantInput()
	fmt.Println("that's end #2.1 ")
}

func outputNumber(size int) {
	for i := 1; i <= size; i++ {
		fmt.Print(i, ",")
	}

	fmt.Println()
}

func outputChar(size int) {
	ch := 'a'
	// var b rune = ch // ch type is rune

	for i := 1; i <= size; i++ {
		fmt.Printf("%c,", ch+rune(i-1))
		//ch++
	}

	fmt.Println()
}

// in one thread, althougth run in concurrently, but only can run one routine in one time
func gorun3_0() {
	go outputNumber(26)
	go outputChar(26)

	wantInput()
}

// in multi thread, run in parrel, can run some routine in one time,
func gorun3_1() {
	trigml()
	gorun3_0()
	//trigsl()
}
