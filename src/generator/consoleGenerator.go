package generator

import (
	"fmt"
	"strings"
)

func GenerateConsole(results map[string]string) {
	for k, v := range results {
		fmt.Println(strings.Repeat("#", 20))
		fmt.Println(k)
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println(v)
		fmt.Println(strings.Repeat("#", 20))
		fmt.Println("")
	}
}
