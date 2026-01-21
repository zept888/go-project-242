package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSize_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := GetSize(tmpDir, Options{})
	require.NoError(t, err)
}

func TestGetSize_NonExistentDir(t *testing.T) {
	err := GetSize("/nonexistent/dir", Options{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to Lstat")
}

func TestGetSize_OneFileDir(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/test.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetSize(tmpDir, Options{})
	assert.NoError(t, err)
}

func TestGetSize_MultipleFilesDir(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/test1.txt", make([]byte, 100), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test2.txt", make([]byte, 200), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test3.txt", make([]byte, 300), 0644))
	err := GetSize(tmpDir, Options{})
	assert.NoError(t, err)
}

func TestGetSize_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/file.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetSize(filepath, Options{})
	require.NoError(t, err)
}

func TestGetSize_WithHumanReadable(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/test.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 1000000), 0644))
	require.NoError(t, GetSize(filepath, Options{HumanReadable: true}))
}

func TestFormatSize_WithoutHum(t *testing.T) {
	result := FormatSize(1024, false)
	assert.Equal(t, "1024B", result)
}

func TestFormatSize_Bytes(t *testing.T) {
	result := FormatSize(500, true)
	assert.Equal(t, "500B", result)
}

func TestFormatSize_Kilobytes(t *testing.T) {
	result := FormatSize(2048, true)
	assert.Equal(t, "2.00KB", result)
}

func TestFormatSize_Megabytes(t *testing.T) {
	result := FormatSize(1048576, true)
	assert.Equal(t, "1.00MB", result)
}

func TestFormatSize_Gigabytes(t *testing.T) {
	result := FormatSize(1073741824, true)
	assert.Equal(t, "1.00GB", result)
}
