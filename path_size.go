package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	HumanReadable bool
	ShowHidden    bool
	Recursive     bool
}

func GetPathSize(path string, opts Options) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("failed to Lstat %s: %w", path, err)
	}

	if !info.IsDir() {
		size := info.Size()
		formatted := FormatSize(size, opts.HumanReadable)
		fmt.Printf("%s\t%s\n", formatted, path)
		return nil
	}

	totalSize, err := CalcDirSize(path, opts)
	if err != nil {
		return err
	}

	formatted := FormatSize(totalSize, opts.HumanReadable)
	fmt.Printf("%s\t%s\n", formatted, path)
	return nil
}

func CalcDirSize(dir string, opts Options) (int64, error) {
	var totalSize int64

	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("failed to read dir %s: %w\n", dir, err)
	}

	for _, file := range files {
		if !opts.ShowHidden && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(dir, file.Name())

		if file.IsDir() {
			subSize, err := CalcDirSize(fullPath, opts)
			if err != nil {
				return 0, err
			}
			totalSize += subSize
			continue
		}

		info, err := file.Info()
		if err != nil {
			return 0, fmt.Errorf("failed to get file info for %s: %w", fullPath, err)
		}
		totalSize += info.Size()
	}
	return totalSize, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
		PB = TB * 1024
		EB = PB * 1024
	)

	switch {
	case size >= EB:
		return fmt.Sprintf("%.2fEB", float64(size)/EB)
	case size >= PB:
		return fmt.Sprintf("%.2fPB", float64(size)/PB)
	case size >= TB:
		return fmt.Sprintf("%.2fTB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.2fGB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2fMB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2fKB", float64(size)/KB)
	default:
		return fmt.Sprintf("%dB", size)
	}
}
