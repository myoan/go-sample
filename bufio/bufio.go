package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func Hoge(str string) string {
	return "Hello, " + str
}

func Fuga(value int) {
	if value > 0 {
		fmt.Println("hoge")
	}
}

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewReader(f)
	for {
		buf, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	}
}
