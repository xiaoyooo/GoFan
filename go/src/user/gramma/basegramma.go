package gramma

import (
	"fmt"
	//"os"
)

func BaseGramma() {
	// baseType()
	// arrayType()
	// sliceType()
	// mapType()
}

func baseType() {
	const i = 0
	const j = 1
	const (
		k = iota
		l = iota
		m = iota
		n = 1
		o = 2
	)

	var say1 string = "abc"
	say2 := say1

	// modify string
	say1InBytes := []byte(say1)
	say1InBytes[0] = 'b'
	say1 = string(say1InBytes)
	// end modify

	fmt.Println(say1, say2)
}

func arrayType() {
	var abytes [2]byte
	abytes[0] = 'c'
	abytes[1] = 'd'

	bbytes := [2]byte{'a', 'b'}
	cbytes := [...]byte{'a', 'b', 'c'}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	dbytes := cbytes
	dbytes[0] = 'z'
	fmt.Println(bbytes, cbytes, easyArray, dbytes)
}

func sliceType() {
	// slice
	var fslice []int // slice, do not know the length
	ar := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fslice = ar[2:5]
	gslice := ar[:5]
	hslice := ar[2:]

	fmt.Println(fslice, gslice, hslice)

	hslice[0] = 100

	// slice is reference type [ptr, len(start,end), maxlen(start, arrend)]
	fmt.Println(fslice, gslice, hslice)
	fmt.Println(len(gslice), cap(gslice))

	// append
	gslice = append(gslice, 101, 102)

	// copy
	var islice []int
	islice = make([]int, len(hslice))
	copy(islice, hslice)

	awone := []int{1, 2, 3, 4, 5, 6}
	copy(awone[2:], awone[4:])

	fmt.Println(awone)

	fmt.Println(fslice, gslice, hslice, islice, ar)
}

// map is dictionary
// map is reference type
// map is not thread safe, must use mutex lock in multi-thread
func mapType() {
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["a"] = 1

	map2 := make(map[string]int)
	map2["b"] = 2

	fmt.Println(map1, map2)
	fmt.Println(map1["a"], map2["b"])

	map3 := map2
	map3["c"] = 3
	fmt.Println(map2, map3)

	// map get return 2 values: value and isContain
	rate, isgot := map2["b"]
	if isgot {
		fmt.Println(rate)
	}
	//map3 = nil
}

func create() {
	// make : for built-in type Creation
	// only for : slice, map, channel
	// make return T, not *T
	// make(T, args)
	//   Call             Type T     Result

	// make(T, n)       slice      slice of type T with length n and capacity n
	// make(T, n, m)    slice      slice of type T with length n and capacity m

	// make(T)          map        map of type T
	// make(T, n)       map        map of type T with initial space for n elements

	// make(T)          channel    synchronous channel of type T
	// make(T, n)       channel    asynchronous channel of type T, buffer size n
	arr := make([]int, 10)
	arr[0] = 1

	// new : for all type Creation
	// new return * T (pointer)

}
