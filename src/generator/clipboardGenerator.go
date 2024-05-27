package generator

import (
	"bytes"
	"github.com/atotto/clipboard"
	"strings"
)

func GenerateClipboard(results map[string]string) {
	var b bytes.Buffer
	for k, v := range results {
		b.WriteString(strings.Repeat("#", 20))
		b.WriteString("\n")
		b.WriteString(k)
		b.WriteString("\n")
		b.WriteString(strings.Repeat("-", 20))
		b.WriteString("\n")
		b.WriteString(v)
		b.WriteString("\n")
		b.WriteString(strings.Repeat("#", 20))
		b.WriteString("\n")
	}

	clipboard.WriteAll(b.String())
}
