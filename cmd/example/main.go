package main

import (
	"flag"
	"fmt"
	lab2 "github.com/merrymike-noname/KPI-lab-2"
	"io"
	"os"
	"errors"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Path to read from file")
	outputFile = flag.String("o", "", "Path to write to file")
)

func main() {
	flag.Parse()

	reader, writer, error := flagHandler() 

	if error != nil {
		fmt.Fprintln(os.Stderr, "Error: ", error)
		os.Exit(1)
	}

	handler := &lab2.ComputeHandler{Reader: reader, Writer: writer}
	err := handler.Compute()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func flagHandler() (io.Reader, io.Writer, error) {
	if (*inputExpression != "" && *inputFile != "") {
		return nil, nil, errors.New("use only one type of input")
	} else if (*inputExpression == "" && *inputFile == "") {
		return nil, nil, errors.New("input data is missing")
	}

	var reader io.Reader 
	var writer io.Writer

	if (*inputExpression != "") {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			return nil, nil, errors.New("error during reading")
		}
		reader = file
	}

	if (*outputFile == "") {
		writer = os.Stdout
	} else {
		file, err := os.Create(*outputFile)
		if err != nil {
			return nil, nil, errors.New("error during writting")
		}
		writer = file
	}
	
	return reader, writer, nil
}