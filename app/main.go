package main

import (
	"fmt"
)

type Planet struct {
	name string
	size uint64 
}

type Program struct {
	program_id uint32
	name string
	member uint32
}

func main() {
	earth := Planet{"Earth", 12742};
	fmt.Println("Earth is a planet named", earth.name, "with a diameter of", earth.size, "km");
}