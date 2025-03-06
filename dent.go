package dent

import (
	"math"
	"regexp"
	"strings"
)

func Indent(text, indent []byte) []byte {
	return []byte(IndentString(string(text), string(indent)))
}

func IndentString(text, indent string) string {
	out := []string{}

	for l := range strings.SplitSeq(text, "\n") {
		out = append(out, indent+l)
	}

	return strings.Join(out, "\n")
}

func Dedent(text []byte) []byte {
	return []byte(DedentString(string(text)))
}

func DedentString(text string) string {
	if text == "" {
		return ""
	}

	if len(strings.Split(text, "\n")) == 1 {
		return strings.TrimLeft(text, " \t")
	}

	out := []string{}

	smallestIndentSize := math.MaxInt
	smallestIndent := ""
	indentRe := regexp.MustCompile(`(?m)(^[ \t]*)[^ \t\n]`)

	for _, match := range indentRe.FindAllStringSubmatch(text, -1) {
		indent := match[1]

		if l := len(indent); l < smallestIndentSize {
			smallestIndent = indent
			smallestIndentSize = l
		}
	}

	for l := range strings.SplitSeq(text, "\n") {
		out = append(out, strings.TrimPrefix(l, smallestIndent))
	}

	return strings.Join(out, "\n")
}
