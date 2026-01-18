package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSize_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := GetSize(tmpDir)
	require.NoError(t, err)
}

func TestGetSize_NonExistentDir(t *testing.T) {
	err := GetSize("/nonexistent/dir")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
}

func TestGetSize_OneFileDir(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/test.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetSize(tmpDir)
	assert.NoError(t, err)
}

func TestGetSize_MultipleFilesDir(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/test1.txt", make([]byte, 100), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test2.txt", make([]byte, 200), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test3.txt", make([]byte, 300), 0644))
	err := GetSize(tmpDir)
	assert.NoError(t, err)
}

func TestGetSize_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/file.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetSize(filepath)
	require.NoError(t, err)
}
