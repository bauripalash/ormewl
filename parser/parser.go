package parser

import (
	"fmt"
	"strings"
)

func MwlTokenize(src string) []string {
	src = strings.ReplaceAll(src, "[", " [ ")
	src = strings.ReplaceAll(src, "]", " ] ")
	src = strings.ReplaceAll(src, "(", " ( ")
	src = strings.ReplaceAll(src, ")", " ) ")
	for value := range strings.FieldsSeq(src) {
		fmt.Printf("['%s']\n", value)
	}
	return strings.Fields(src)
}
