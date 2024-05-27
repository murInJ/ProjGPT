package generator

import "fmt"

func GenerateProj(parseResults map[string]string, generateConsole bool, generateClipboard bool) error {
	if generateConsole {
		GenerateConsole(parseResults)
	}
	if generateClipboard {
		GenerateClipboard(parseResults)
		fmt.Println("generate to clipboard.")
	}
	return nil
}
