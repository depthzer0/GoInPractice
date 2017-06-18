package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish")
	flag.BoolVar(&spanish, "s", true, "Use Spanish")
}

func main() {
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(spanish)
}
