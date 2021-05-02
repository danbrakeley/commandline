package commandline

import (
	"testing"
)

func Test_ParseCommandLine(t *testing.T) {
	cases := []struct {
		ShouldSucceed bool
		Name          string
		In            string
		Out           []string
	}{
		{true, "one arg", `foo`, []string{"foo"}},
		{true, "two args", `foo bar`, []string{"foo", "bar"}},
		{true, "one arg with space", `foo\ bar`, []string{"foo bar"}},
		{true, "string with quotes", `foo "bar bar bar"`, []string{"foo", "bar bar bar"}},
		{true, "space prefix", ` space`, []string{"space"}},
		{true, "tab prefix", "\ttab", []string{"tab"}},
		{true, "space suffix", `baz `, []string{"baz"}},
		{true, "space prefix and suffix", ` 88 `, []string{"88"}},
		{true, "three args with lots of whitespace", ` foo  -v     bar `, []string{"foo", "-v", "bar"}},
		{true, "single quote inside double quotes", `bar --commit "it's done"`, []string{"bar", "--commit", "it's done"}},
		{true, "quoted quotes", `"'"`, []string{"'"}},

		{false, "unmatched double quote", `"`, nil},
		{false, "unmatched double quote followed by matched single quotes", `"''`, nil},
		{false, "unmatched singled quote followed by double quotes", `'""`, nil},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			if !tc.ShouldSucceed {
				_, err := Parse(tc.In)
				if err == nil {
					t.Errorf("expected Parse(`%s`) to fail, but it succeeded", tc.In)
				}
				return
			}

			expected := tc.Out
			actual, err := Parse(tc.In)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			areEqual := len(expected) == len(actual)
			if areEqual {
				for i := range expected {
					if actual[i] != expected[i] {
						areEqual = false
						break
					}
				}
			}
			if !areEqual {
				t.Errorf("expected %v, actual %v", expected, actual)
			}

		})
	}
}
