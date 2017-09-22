package common

import (
	"flag"
	"fmt"
)

// PrintFlags 输出所有已解析的 flag
func PrintFlags() {

	// 防止用户忘记解析
	if !flag.Parsed() {
		flag.Parse()
	}

	visitor := func(a *flag.Flag) {
		fmt.Println("flag =", a.Name, "\t", " value =", a.Value, "\t", a.Usage)
	}
	fmt.Println("=-------=")
	fmt.Println("| flags |")
	fmt.Println("=-------=")
	flag.VisitAll(visitor)
}

// GetFlagByName 通过 flag name 获取 value
func GetFlagByName(name string) (value string, err error) {
	flag := flag.Lookup(name)
	if flag == nil {
		return "", fmt.Errorf("Can't find option[\"%s\"]", name)
	}
	return flag.Value.String(), nil
}
