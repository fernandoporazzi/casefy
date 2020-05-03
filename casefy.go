package casefy

import (
	"log"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// replaceNonAlphaNumericChars returns a copy of s, replacing matches of the Regexp
// with the replacement string repl.
func replaceNonAlphaNumericChars(s, repl string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, repl)
}

func applyKebabOrSnakeCase(input, repl string) string {
	input = strings.ToLower(input)

	output := replaceNonAlphaNumericChars(input, repl)

	output = strings.TrimRight(output, repl)
	output = strings.TrimLeft(output, repl)

	return output
}

// ToPascalCase returns input in Pascal Case format
//
// "my nice input" becomes "MyNiceInput"
func ToPascalCase(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}

	input = strings.ToLower(input)
	input = strings.Title(input)

	return replaceNonAlphaNumericChars(input, "")
}

// ToCamelCase returns input in Camel Case format
//
// "my nice input" becomes "myNiceInput"
func ToCamelCase(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}

	input = strings.ToLower(input)
	input = strings.Title(input)
	input = replaceNonAlphaNumericChars(input, "")

	rune, size := utf8.DecodeRuneInString(input)
	return string(unicode.ToLower(rune)) + input[size:]
}

// ToKebabCase returs input in Kebab Case Format.
//
// "my nice input" becomes "my-nice-input"
func ToKebabCase(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}

	return applyKebabOrSnakeCase(input, "-")
}

// ToSnakeCase returs input in Snake Case Format.
//
// "my nice input" becomes "my_nice_input"
func ToSnakeCase(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}

	return applyKebabOrSnakeCase(input, "_")
}
