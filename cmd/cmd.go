package main

import (
	"fmt"
	"os"
)

func main() {
	//takes every argument
	command := os.Args
	//	./cmd new file ---prints the whole argument
	fmt.Println(command)

	secondArgument := os.Args[1:]
	//	./cmd new file ---skips the first part of the argument and takes the rest
	fmt.Println(secondArgument)

}
