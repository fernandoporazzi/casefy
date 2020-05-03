package main

import (
	"fmt"

	"github.com/fernandoporazzi/casefy"
)

func main() {
	fmt.Println(casefy.ToPascalCase("foo bar and john doe"))                    // => FooBarAndJohnDoe
	fmt.Println(casefy.ToCamelCase("hello amsterdam, it is a sunny day"))       // => helloAmsterdamItIsASunnyDay
	fmt.Println(casefy.ToKebabCase("--foo bar 1 and john doe"))                 // => foo-bar-1-and-john-doe
	fmt.Println(casefy.ToSnakeCase("a very nice string, should be snake case")) // => a_very_nice_string_should_be_snake_case
}
