package main

import (
	"fmt"
	"github.com/lebrancconvas/Planetarium/shape"
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
	program01 := shape.Circle{Radius: 10};
	fmt.Println(program01.Circumstance());
}