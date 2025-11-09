package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

const (
	flagLines     = "lines"
	flagWords     = "words"
	flagChars     = "chars"
	flagBytes     = "bytes"
	flagMaxLength = "max-line-length"
)

func main() {
	app := &cli.App{
		Name:  "wc",
		Usage: "print newline, word, and byte counts for each file",
		UsageText: `wc [OPTIONS] [FILE...]

   Print newline, word, and byte counts for each FILE, and a total line if
   more than one FILE is specified. A word is a non-zero-length sequence of
   characters delimited by white space.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagLines,
				Aliases: []string{"l"},
				Usage:   "print the newline counts",
			},
			&cli.BoolFlag{
				Name:    flagWords,
				Aliases: []string{"w"},
				Usage:   "print the word counts",
			},
			&cli.BoolFlag{
				Name:    flagChars,
				Aliases: []string{"m"},
				Usage:   "print the character counts",
			},
			&cli.BoolFlag{
				Name:    flagBytes,
				Aliases: []string{"c"},
				Usage:   "print the byte counts",
			},
			&cli.BoolFlag{
				Name:    flagMaxLength,
				Aliases: []string{"L"},
				Usage:   "print the maximum display width",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "wc: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagLines) {
		params = append(params, Lines)
	}
	if c.Bool(flagWords) {
		params = append(params, Words)
	}
	if c.Bool(flagChars) {
		params = append(params, Chars)
	}
	if c.Bool(flagBytes) {
		params = append(params, Bytes)
	}
	if c.Bool(flagMaxLength) {
		params = append(params, MaxLength)
	}

	// Create and execute the wc command
	cmd := Wc(params...)
	return gloo.Run(cmd)
}
