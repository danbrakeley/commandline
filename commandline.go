package commandline

import (
	"fmt"
)

const (
	sStart = iota
	sQuotes
	sArg
)

// Parse takes a string and makes an array of the command and its arguments.
// Respects escaped quotes, spaces, and other characters.
// based on https://stackoverflow.com/questions/34118732/parse-a-command-line-string-into-flags-and-arguments-in-golang
func Parse(command string) ([]string, error) {
	var args []string
	state := sStart
	current := ""
	quote := ""
	escapeNext := false
	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == sQuotes {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = sStart
			}
			continue
		}

		if escapeNext {
			current += string(c)
			escapeNext = false
			continue
		}

		if c == '\\' {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = sQuotes
			quote = string(c)
			continue
		}

		if state == sArg {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = sStart
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = sArg
			current += string(c)
		}
	}

	if state == sQuotes {
		return []string{}, fmt.Errorf("unclosed quote in command line: %s", command)
	}

	if len(current) > 0 {
		args = append(args, current)
	}

	return args, nil
}
