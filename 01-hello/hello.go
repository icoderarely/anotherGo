// Package hello provides: hello, like really
package hello

import (
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	// return fmt.Sprintf("Hello, %s!", name)
	return "Hello, " + strings.Join(names, ", ") + "!"
}
