package lab2

import(
	"io"
	"bytes"
	"strings"
	"fmt"
)

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var buffer bytes.Buffer

	_, readErr := buffer.ReadFrom(ch.Reader)
	if readErr != nil {
		return readErr
	}

	correctInput := strings.TrimRight(buffer.String(), "\n")

	result, inputErr := CalculatePostfix(correctInput) 
	if inputErr != nil {
		return inputErr
	}

	fmt.Fprintln(ch.Writer, result)
	return nil
}