package main

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeGroups(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, group := range groups {
		cp := append([]string(nil), group...)
		sort.Strings(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		if len(out[i]) != len(out[j]) {
			return len(out[i]) < len(out[j])
		}
		for k := range out[i] {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

type groupTestCase struct {
	name  string
	input []string
	want  [][]string
}

func TestGroupAnagrams(t *testing.T) {
	tests := []groupTestCase{
		{
			name:  "example mix",
			input: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want:  [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name:  "single string",
			input: []string{"x"},
			want:  [][]string{{"x"}},
		},
		{
			name:  "single empty string",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "multiple empty strings",
			input: []string{"", ""},
			want:  [][]string{{"", ""}},
		},
		{
			name:  "duplicates",
			input: []string{"a", "a", "b"},
			want:  [][]string{{"a", "a"}, {"b"}},
		},
		{
			name:  "all anagrams",
			input: []string{"abc", "bca", "cab", "acb"},
			want:  [][]string{{"abc", "bca", "cab", "acb"}},
		},
		{
			name:  "no anagrams",
			input: []string{"a", "b", "c"},
			want:  [][]string{{"a"}, {"b"}, {"c"}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := normalizeGroups(tc.want)

			got := normalizeGroups(groupAnagrams(tc.input))
			if !reflect.DeepEqual(got, want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tc.input, got, want)
			}

			gotBetter := normalizeGroups(BetterApproach(tc.input))
			if !reflect.DeepEqual(gotBetter, want) {
				t.Errorf("BetterApproach(%v) = %v, want %v", tc.input, gotBetter, want)
			}
		})
	}
}
