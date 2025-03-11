package main

import (
	"fmt"
	"strings"
)

func rangeIV() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Printf("%d,", pow[i])
	}
	for _, value := range pow {
		fmt.Printf("\n%d", value)
	}
}

func rangeSlice() {
	x := []int{0, 2, 4, 6, 8, 10}
	for i, v := range x {
		fmt.Printf("2 ** %d = %d\n", i, v)
	}
}

func ticTacToe() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

var arr = [2]string{"Hello", "World"}

var primes = [8]int{1, 2, 3, 5, 7, 11, 13, 19}
var slicePrime1 []int = primes[1:6]
var slicePrime2 []int = primes[5:8]

// slices can be created without array also
var justSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// slice defaults
var justSliceStart = justSlice[:3]
var justSliceEnd = justSlice[3:]

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makingSlice() {
	// if capacity is not given, it is taken same as length
	a := make([]int, 5)
	printMakingSlice("a", a)
	b := make([]int, 0, 5)
	printMakingSlice("b", b)
	c := b[:2]
	printMakingSlice("c", c)
	d := c[2:5]
	printMakingSlice("d", d)
}

func printMakingSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
