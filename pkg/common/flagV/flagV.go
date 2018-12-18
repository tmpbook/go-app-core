package flagV

import (
	"flag"
	"fmt"
)

// PrintFlags print all parsed flags
func PrintFlags() {

	// prevent users from forgetting
	if !flag.Parsed() {
		flag.Parse()
	}

	visitor := func(a *flag.Flag) {
		fmt.Println("flag =", a.Name, "\t", " value =", a.Value, "\t", "default =", a.DefValue, "\t", a.Usage)
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println(" flags:")
	fmt.Println("------------------------------------------------------------------------------")
	flag.VisitAll(visitor)
}
