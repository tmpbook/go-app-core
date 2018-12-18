package main

import (
	"flag"
	"fmt"

	"github.com/tmpbook/go-app-core/pkg/common/flagV"
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
	flagV.PrintFlags()
	fmt.Printf("%s:%s %t", hostname, port, open)
}
