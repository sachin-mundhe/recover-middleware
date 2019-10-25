package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCases struct {
	name   string
	input  string
	output string
}

func TestGenerateLinks(t *testing.T) {

	var tcs []testCases = []testCases{
		{name: "No input value"},
		{name: "Valid Input",
			input:  "This is valid stack strace\n\t/usr/local/go/bin:45",
			output: "This is valid stack strace\n\t<a href=\"/sourcecode/?line=45&path=%2Fusr%2Flocal%2Fgo%2Fbin\">/usr/local/go/bin</a>:45",
		},
		{name: "Invalid Line Number",
			input:  "This is valid stack strace\n\t/usr/local/go/bin: 45",
			output: "This is valid stack strace\n\t<a href=\"/sourcecode/?line=&path=%2Fusr%2Flocal%2Fgo%2Fbin\">/usr/local/go/bin</a>: 45",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(ts *testing.T) {
			// fmt.Println("Link Input:", tc.input)

			link := GenerateLinks(tc.input)
			// fmt.Println("Link Output:", link)
			assert.EqualValues(ts, tc.output, link)
		})
	}

}
