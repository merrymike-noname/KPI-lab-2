package main

import (
	"flag"
	"fmt"
	lab2 "github.com/merrymike-noname/KPI-lab-2"
	"io"
	"os"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Path to read from file")
	outputFile = flag.String("o", "", "Path to write to file")
)

func main() {
	flag.Parse()

	var reader io.Reader
	var writer io.Writer 
	var error error  

	if error != nil {
		fmt.Fprintln(os.Stderr, "Error: ", error)
		os.Exit(1)
	}

	handler := &lab2.ComputeHandler{Reader: reader,Writer: writer}
	err := handler.Compute()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}
