# commandline

## Overview

Parses single string commandlines into command/arguments for use with the Go standard library os/exec.

Examples (see [commandline_test.go](https://github.com/danbrakeley/commandline/blob/main/commandline_test.go) for more):

input | output
--- | ---
`foo bar baz` | `[]string{"foo", "bar", "baz"}`
`foo bar 'baz bif'` | `[]string{"foo", "bar", "baz bif"}`
`bar --commit "it's done"` | `[]string{"bar", "--commit", "it's done"}`

## License

See `LICENSE.txt`.

License may change in the future, so always read the license file included with the specific version you want to use.

While it isn't required, if you have any improvements, I'd love to see them and possibly incorporate them,
so please drop me a note or send me a pull reqeust!
