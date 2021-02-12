package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ExtractHeadings(t *testing.T) {
	before := "&lt;br/&gt;&lt;br/&gt;&lt;b&gt;"
	after := "&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;"
	text := fmt.Sprintf("something %sCharacters%s something %sMonsters%s something", before, after, before, after)
	expected := []string{"Characters", "Monsters"}
	if actual := extractHeadings(text); !reflect.DeepEqual(expected, actual) {
		t.Errorf("got %q; expected: %q", actual, expected)
	}
	if actual := extractHeadings(""); !reflect.DeepEqual([]string{}, actual) {
		t.Errorf("got %q; expected: %q", actual, expected)
	}
}
