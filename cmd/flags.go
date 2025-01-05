package main

import (
	"flag"
	"fmt"
)

func main() {
	words := flag.String("ls", "list", "getting all lists")

	flag.Parse()
	fmt.Println(*words)
}
