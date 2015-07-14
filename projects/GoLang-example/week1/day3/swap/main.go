package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func RotateRight(xs ...*int) {
	first := *xs[0]
	for i := 1; i < len(xs); i++ {
		*xs[i-1] = *xs[i]
	}
	*xs[len(xs)-1] = first
}

func main() {
	x, y := 5, 6
	swap(&x, &y)
	fmt.Println(x, y)
	a, b, c, d, e, f, g, h := 1, 2, 3, 4, 5, 6, 7, 8
	RotateRight(&a, &b, &c, &d, &e, &f, &g, &h)
	fmt.Println(a, b, c, d, e, f, g, h)
}
