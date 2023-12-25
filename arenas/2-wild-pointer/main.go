package main

import (
	"arena"
	"fmt"
	"unsafe"
)

func write() {
	mem := arena.NewArena()

	s := arena.MakeSlice[int](mem, 10, 10)
	fmt.Printf("array ptr: %p\n", unsafe.Pointer(&s[0]))

	mem.Free()

	s[0] = 42
}

func read() {
	mem := arena.NewArena()

	s := arena.MakeSlice[int](mem, 10, 10)
	fmt.Printf("array ptr: %p\n", unsafe.Pointer(&s[0]))

	mem.Free()

	s = append(s, 42)
}

func main() {
	// write()
	read()
}
