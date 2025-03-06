package dent_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/nobe4/dent.go"
)

func TestIndent(t *testing.T) {
	tests := []struct {
		text   string
		indent string
		want   string
	}{
		{},
		{
			text:   "",
			indent: "  ",
			want:   "  ",
		},
		{
			text:   "a",
			indent: "  ",
			want:   "  a",
		},
		{
			text: "a\n" +
				"b",
			indent: "  ",
			want: "  a\n" +
				"  b",
		},
		{
			text: "a\n" +
				"b\n",
			indent: "xxx",
			want: "xxxa\n" +
				"xxxb\n" +
				"xxx",
		},
	}

	for _, test := range tests {
		t.Run("Byte", func(t *testing.T) {
			got := dent.Indent([]byte(test.text), []byte(test.indent))
			if !bytes.Equal(got, []byte(test.want)) {
				t.Fatalf("\nwant %q\ngot  %q", test.want, got)
			}
		})

		t.Run("String", func(t *testing.T) {
			got := dent.IndentString(test.text, test.indent)
			if !strings.EqualFold(got, test.want) {
				t.Fatalf("\nwant %q\ngot  %q", test.want, got)
			}
		})
	}
}

func TestDedent(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{},
		{
			text: "   ",
			want: "",
		},

		{
			text: "a",
			want: "a",
		},

		{
			text: "a\n" +
				"b",
			want: "a\n" +
				"b",
		},

		{
			text: "a\n" +
				"b\n",
			want: "a\n" +
				"b\n" +
				"",
		},

		{
			text: "  a\n" +
				"b",
			want: "  a\n" +
				"b",
		},

		{
			text: `
    a
        b
            c`,
			want: `
a
    b
        c`,
		},

		{
			text: `
	a
	    b
	        c
`,
			want: `
a
    b
        c
`,
		},

		{
			text: `
	a
	  
	    b
	  
	        c
`,
			want: `
a
  
    b
  
        c
`,
		},
		{
			text: `
	a
  
	    b
  
	        c
`,
			want: `
a
  
    b
  
        c
`,
		},
	}

	for _, test := range tests {
		t.Run("Byte", func(t *testing.T) {
			got := dent.Dedent([]byte(test.text))
			if !bytes.Equal(got, []byte(test.want)) {
				t.Fatalf("\nwant %q\ngot  %q", test.want, got)
			}
		})

		t.Run("String", func(t *testing.T) {
			got := dent.DedentString(test.text)
			if !strings.EqualFold(got, test.want) {
				t.Fatalf("\nwant %q\ngot  %q", test.want, got)
			}
		})
	}
}

func ExampleIndent() {
	text := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
rutrum dui.
`
	fmt.Println(dent.IndentString(text, "  "))
	// Output:
	//   Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
	//   facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
	//   accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
	//   rutrum dui.
}

func ExampleDedent() {
	text := `
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
    facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
    accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
    rutrum dui.
`
	fmt.Println(dent.DedentString(text))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
	// facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
	// accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
	// rutrum dui.
}

func BenchmarkIndent(b *testing.B) {
	text := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
rutrum dui.
`
	for b.Loop() {
		dent.IndentString(text, "  ")
	}
}

func BenchmarkDedent(b *testing.B) {
	text := `
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
    facilisis tempus fringilla. Nullam nunc massa, rhoncus nec enim in, egestas
    accumsan odio. Nullam ligula felis, suscipit consequat sapien et, efficitur
    rutrum dui.
`
	for b.Loop() {
		dent.DedentString(text)
	}
}
