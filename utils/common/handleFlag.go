package common

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
		fmt.Println("flag =", a.Name, "\t", " value =", a.Value, "\t", a.Usage)
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println(" flags:")
	fmt.Println("------------------------------------------------------------------------------")
	flag.VisitAll(visitor)
}

// GetFlagByName get value by flag name
func GetFlagByName(name string) (value string, err error) {
	flag := flag.Lookup(name)
	if flag == nil {
		return "", fmt.Errorf("Can't find flag[\"%s\"]", name)
	}
	return flag.Value.String(), nil
}
