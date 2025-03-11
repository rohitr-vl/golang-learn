package main

import (
	"fmt"
	"math"
	"strings"
)

// closure
func fibonacci() func() int {
	i, j := 0, 1
    return func() int {
		k:=j
		j += i
		i = k
        return i + j
    }
}

func printFibonnaci(upperLimit int) {
	f := fibonacci()
	for i := 0; i < upperLimit; i++ {
		fmt.Println(f())
	}
	fmt.Println("")
}

func intSeq() func() int {
    i := 0
    return func() int {
        i += 5
        return i
    }
}

func closureExample() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println("")
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcAsParam() {
	// hypotenuse = sqrt((a2 + b2))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func WordCount(s string) map[string]int {
	stringSlice := strings.Fields(s)
	wCount := make(map[string]int)
	for cnt := range len(stringSlice) {
		wCount[stringSlice[cnt]] += strings.Count(s, stringSlice[cnt])
	}
	return wCount
}

type Vortex struct {
	Lat, Long float64
}

// 1st way of initializing a map
var mVortex = map[string]Vortex{
	"Google": {
		37.42202, -122.08408,
	},
}

func useMaps() {
	// mVortex = make(map[string]Vortex)
	//2nd way of initializing a map
	mVortex["Bell Labs"] = Vortex{
		40.68433, -74.39967,
	}
	fmt.Println("")
	fmt.Println(mVortex)

	// Test that a key is present with a two-value assignment:
	val, ok := mVortex["Bell Labs"]
	fmt.Println("The value:", val, "Present?", ok)

    m := make(map[string]int)
	m["k1"] = 7
    m["k2"] = 13
	fmt.Println("map:", m)
	val1 := m["k1"]
	fmt.Println("val1: ", val1)
	val2 := m["k2"]
	fmt.Println("val2: ", val2)
	val3 := m["k3"]
    fmt.Println("val3:", val3)
	fmt.Println("len of m:", len(m))
	delete(m, "k2")
    fmt.Println("After delete of k2, m:", m, " len of m:", len(m))
	m["k2"] = 15
	m["k3"] = 550
	fmt.Println("After alloc of k2, m:", m, " len of m:", len(m))
	clear(m)
	fmt.Println("After clear of m, m:", m, " len of m:", len(m))
}