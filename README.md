# casefy

Convert a string to camelCase, snake_case, kebab-case or PascalCase

## Install

```
$ go get github.com/fernandoporazzi/casefy
```

## Usage

```go
casefy.ToPascalCase("foo bar and john doe")
//=> FooBarAndJohnDoe

casefy.ToCamelCase("hello amsterdam, it is a sunny day")
//=> helloAmsterdamItIsASunnyDay

casefy.ToKebabCase("--foo bar 1 and john doe")
//=> foo-bar-1-and-john-doe

casefy.ToSnakeCase("a very nice string, should be snake case")
//=> a_very_nice_string_should_be_snake_case
```
