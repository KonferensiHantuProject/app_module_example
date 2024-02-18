package main

import (
	"fmt"

	// V1
	// go_module_example "github.com/MuhammadNazhimMaulana/go_module_example"

	// V2
	go_module_example "github.com/MuhammadNazhimMaulana/go_module_example/v2"
)

func main() {
	fmt.Println(go_module_example.SayHello("Anjas"))
}
