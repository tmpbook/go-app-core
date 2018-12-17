package main

import (
	"flag"
	"fmt"
)

var (
	hostname string
	port     string
	open     bool
)

func init() {
	flag.StringVar(&hostname, "h", "localhost", "hostname")
	flag.StringVar(&port, "p", "8888", "port")
	flag.BoolVar(&open, "o", false, "open the page")

	flag.Parse()
}

func main() {
	fmt.Printf("%s:%s %t", hostname, port, open)
}

// go run main.go -h=8.8.8.8 -p=80 -o=true
