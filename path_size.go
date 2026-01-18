package code

import (
	"fmt"
	"os"
)

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
