package main

import (
	"reflect"
	"strings"
	"testing"
)

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return reflect.DeepEqual(a, b)
}

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{name: "example", input: []string{"Hello", "World"}},
		{name: "single empty", input: []string{""}},
		{name: "empty list", input: []string{}},
		{name: "hash characters", input: []string{"#", "##", "a#b"}},
		{name: "embedded delimiters", input: []string{"a,b", "c;d", "1#2#3"}},
		{name: "empty mixed", input: []string{"", "abc", "", "def"}},
		{name: "spaces", input: []string{" hello ", "world"}},
		{name: "unicode", input: []string{"héllo", "世界", "😀"}},
		{name: "long string", input: []string{strings.Repeat("x", 200)}},
	}

	s := &Solution{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encoded := s.Encode(tc.input)
			got := s.Decode(encoded)
			if !equalStrings(got, tc.input) {
				t.Errorf("Decode(Encode(%v)) = %v, want %v", tc.input, got, tc.input)
			}
		})
	}
}

func TestEncodeFormat(t *testing.T) {
	s := &Solution{}
	got := s.Encode([]string{"Hello", "World"})
	want := "5#Hello5#World"
	if got != want {
		t.Errorf("Encode() = %q, want %q", got, want)
	}
}
