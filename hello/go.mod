module scottschubert.dev/hello

go 1.23.0

replace scottschubert.dev/greetings => ../greetings

require (
	scottschubert.dev/greetings v0.0.0-00010101000000-000000000000
	scottschubert.dev/some_other_package v0.0.0-00010101000000-000000000000
)

replace scottschubert.dev/some_other_package => ../package
