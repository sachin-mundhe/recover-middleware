package links

import (
	"net/url"
	"strings"
)

//GenerateLinks This function accepts string as parameter and make links clickable
func GenerateLinks(rawData string) string {

	lines := strings.Split(rawData, "\n")

	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""

		for i, v := range line {
			if v == ':' {
				file = line[1:i]
				break
			}
		}

		//Line Number
		var lineStr strings.Builder
		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}
			lineStr.WriteByte(line[i])
		}

		v := url.Values{}
		v.Set("line", lineStr.String())
		v.Set("path", file)

		lines[li] = "\t<a href=\"/sourcecode/?" + v.Encode() + "\">" + file + "</a>" + line[len(file)+1:]
	}

	output := strings.Join(lines, "\n")
	return output
}
