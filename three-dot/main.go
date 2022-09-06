package main

import "fmt"

func populate(v ...interface{}) {
	fmt.Printf("len: %d\n", len(v))
}

func execContext(v ...interface{}) {
	populate(v)
}

func execContext2(v ...interface{}) {
	populate(v...)
}

func main() {
	ar := []interface{}{1, 2, 3}
	execContext(ar...)
	execContext2(ar...)
}
