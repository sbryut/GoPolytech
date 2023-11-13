//go:build pro
// +build pro

package main

func main() {
	multip := 1
	for i := 1; i < 8; i++ {
		multip *= i
	}
	println("Multiplication = ", multip)
}
