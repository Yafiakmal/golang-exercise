package main

import (
	"fmt"
	"flag"
	"github.com/gomarkdown/markdown"
	"os"
)

const SIZEMAX int64 = 104857600

func main() {
	inputPath, outputPath := getArgs()
	checkStat(inputPath)
	mdToHtml(inputPath, outputPath)
}

func getArgs() (string, string) {
	inputPath := flag.String("input", "", "Path file markdown")
	outputPath := flag.String("output", "output/output.html", "Path file HTML output")
	flag.Parse()

	if *inputPath == "" {
		fmt.Fprintln(os.Stderr, "\n\nError: -input flag must be provided")
		flag.Usage()
		os.Exit(1)
	}

	return *inputPath, *outputPath
}

func checkStat(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please check your file.md path\n   ", err)
		os.Exit(1)
	}
	size := info.Size()
	if size >= SIZEMAX {
		fmt.Fprintf(os.Stderr, "maximum size are : %d mb. but yours: %d\n", SIZEMAX/1024/1024, size/1024/1024 )
		os.Exit(1)
	}
}

func mdToHtml(input, output string) {
	inputData, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	result := markdown.ToHTML(inputData, nil, nil)

	err = os.WriteFile(output, result, 0644)
	if err != nil {
		panic(err)
	}
}

// func (args) returnType {
	
// }

/*

with closing
*, **, ***, _, __, ___, (), [], <>, ``


no closing
!, #, >, *, -, +,

rule
1. tidak ada endl lebih dari 2
2. buffer final harus berakhiran closing, endl, atau eof
*/
