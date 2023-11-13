//go:build !pro
// +build !pro

package main

func main() {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += i
	}
	println("Sum = ", sum)
}
