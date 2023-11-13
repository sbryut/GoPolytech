package main

import (
	"embed"
	"fmt"
)

//go:embed read.txt
var readedFile string

//go:embed read.txt
var folder embed.FS

func main() {
	fmt.Println("Reading file in last task: ")
	fmt.Println(readedFile)
}
