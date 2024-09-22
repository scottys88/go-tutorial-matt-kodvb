package main

import (
	"fmt"

	"scottschubert.dev/greetings"
	"scottschubert.dev/some_other_package"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	some_other_package := some_other_package.SomeOtherThing()
	fmt.Println(message)
	fmt.Println(some_other_package)
}
