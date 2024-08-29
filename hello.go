package hello

import (
	"fmt"
	"strings"
)

func Say(names []string) string {
	fmt.Println(names)
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello " + strings.Join(names, ", ") + "!"
}
