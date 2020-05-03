package casefy

import "testing"

func TestToPascalCase(t *testing.T) {
	var entries = []struct {
		input string
		want  string
	}{
		{"--foo --bar", "FooBar"},
		{"foo-BaR ", "FooBar"},
		{"this is an entry", "ThisIsAnEntry"},
		{"an entry with a number 1", "AnEntryWithANumber1"},
		{"a @ very _ ! - dirty ```` = entry", "AVeryDirtyEntry"},
	}

	for _, entry := range entries {
		got := ToPascalCase(entry.input)
		if got != entry.want {
			t.Errorf("Expected result to be %v, but got %v", entry.want, got)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	var entries = []struct {
		input string
		want  string
	}{
		{"--foo --bar", "fooBar"},
		{"foo-BaR ", "fooBar"},
		{"this is an entry", "thisIsAnEntry"},
		{"an entry with a number 1", "anEntryWithANumber1"},
		{"a @ very _ ! - dirty ```` = entry", "aVeryDirtyEntry"},
	}

	for _, entry := range entries {
		got := ToCamelCase(entry.input)
		if got != entry.want {
			t.Errorf("Expected result to be %v, but got %v", entry.want, got)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	var entries = []struct {
		input string
		want  string
	}{
		{"--foo --bar", "foo_bar"},
		{"foo-BaR ", "foo_bar"},
		{"this is an entry", "this_is_an_entry"},
		{"an entry with a number 1", "an_entry_with_a_number_1"},
		{"a @ very _ ! - dirty ```` = entry", "a_very_dirty_entry"},
	}

	for _, entry := range entries {
		got := ToSnakeCase(entry.input)
		if got != entry.want {
			t.Errorf("Expected result to be %v, but got %v", entry.want, got)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	var entries = []struct {
		input string
		want  string
	}{
		{"--foo --bar", "foo-bar"},
		{"foo-BaR ", "foo-bar"},
		{"this is an entry", "this-is-an-entry"},
		{"an entry with a number 1", "an-entry-with-a-number-1"},
		{"a @ very _ ! - dirty ```` = entry", "a-very-dirty-entry"},
	}

	for _, entry := range entries {
		got := ToKebabCase(entry.input)
		if got != entry.want {
			t.Errorf("Expected result to be %v, but got %v", entry.want, got)
		}
	}
}
