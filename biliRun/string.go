package main

import (
	"fmt"
	"strings"
)

func makeEfficientString(str string) string {
	var builder strings.Builder
	builder.Grow(20)
	builder.WriteString(str)
	builder.WriteString(":")
	builder.WriteString("case")
	return builder.String()
}

func makeFmtString(str string) string {
	return fmt.Sprintf("%s:case", str)
}
