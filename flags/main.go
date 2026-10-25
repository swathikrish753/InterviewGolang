package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "swathi", "specify the name to greet")
	age := flag.Int("age", 2, "specify the age")
	flag.Parse()
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
