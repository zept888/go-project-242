package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3" // imports as package "cli"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Usage:   "human-readable sizes (auto-select unit) (default: false)",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    "all",
				Usage:   "include hidden files and directories (default: false)",
				Aliases: []string{"a"},
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Usage:   "recursive size of directories (default: false)",
				Aliases: []string{"r"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			if path == "" {
				fmt.Println(".hexlet-path-size <path required>")
				return nil
			}
			opts := code.Options{
				HumanReadable: cmd.Bool("human"),
				ShowHidden:    cmd.Bool("all"),
			}
			return code.GetSize(path, opts)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
