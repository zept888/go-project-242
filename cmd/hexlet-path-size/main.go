package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3" // imports as package "cli"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			if path == "" {
				return fmt.Errorf("path is required")
			}
			return GetSize(path)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func GetSize(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("failed to Lstat %s: %w", path, err)
	}

	if !info.IsDir() {
		size := info.Size()
		fmt.Printf("%d\t%s\n", size, path)
	} else {
		var totalSize int64
		files, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("failed to read dir %s: %w\n", path, err)
		}
		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				return fmt.Errorf("failed to get file info for %s: %w", file.Name(), err)
			}
			totalSize += info.Size()

		}
		fmt.Printf("%d\t%s\n", totalSize, path)
	}
	return nil
}
