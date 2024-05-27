package main

import (
	"ProjGPT/src/generator"
	"ProjGPT/src/parser"
	"ProjGPT/src/utils"
	"flag"
)

func main() {
	rootPtr := flag.String("root", ".", "proj`s root path")
	parseStructPtr := flag.Bool("p_struct", true, "true if parse project`s struct")
	generateConsolePtr := flag.Bool("g_console", false, "true if generate results and echo in console")
	generateClipboardPtr := flag.Bool("g_clipboard", true, "generate to clipboard")
	flag.Parse()

	utils.Check_root(rootPtr)

	res, err := parser.ParseProj(*rootPtr, *parseStructPtr)
	if err != nil {
		panic(err)
	}

	err = generator.GenerateProj(res, *generateConsolePtr, *generateClipboardPtr)
	if err != nil {
		panic(err)
	}
}
