package main

import "fmt"

func halfer(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func makeOddGenerator(x int) int {
	if x&2 == 0 {
		return (x + 1)
	} else {
		return x
	}
}

func fibSeq(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return (fibSeq(n-1) + fibSeq(n-2))
	}
}

func reverse(xs []int) []int {
	ys := make([]int, len(xs))
	for k, v := range xs {
		ys[len(xs)-(k+1)] = v
	}
	return ys
}

func main() {
	var num int
	//-------------------------------------------------
	// fmt.Println("Enter number to be cut: ")
	// fmt.Scanln(&num)
	// fmt.Println(halfer(num))
	//-------------------------------------------------
	// fmt.Println("Enter number to be made odd: ")
	// fmt.Scanln(&num)
	// fmt.Println(makeOddGenerator(num))
	//-------------------------------------------------
	// fmt.Println("Enter nth term for fibinacci: ")
	// fmt.Scanln(&num)
	// fmt.Println(fibSeq(num))
	//-------------------------------------------------
	var array []int
	fmt.Println("Enter list of numbers to reverse: ")
	fmt.Scanln(&num)
	fmt.Println(reverse(num))
}
