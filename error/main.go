package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("hogehoge")

	fmt.Printf("err: %v\n", err)

	err = errors.New("fugafuga")
	fmt.Printf("err: %v\n", err)
}
