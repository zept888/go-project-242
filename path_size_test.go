package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSize_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := GetPathSize(tmpDir, Options{})
	require.NoError(t, err)
}

func TestGetSize_NonExistentDir(t *testing.T) {
	err := GetPathSize("/nonexistent/dir", Options{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to Lstat")
}

func TestGetSize_OneFileDir(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/test.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetPathSize(tmpDir, Options{})
	assert.NoError(t, err)
}

func TestGetSize_MultipleFilesDir(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/test1.txt", make([]byte, 100), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test2.txt", make([]byte, 200), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/test3.txt", make([]byte, 300), 0644))
	err := GetPathSize(tmpDir, Options{})
	assert.NoError(t, err)
}

func TestGetSize_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/file.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 100), 0644))
	err := GetPathSize(filepath, Options{})
	require.NoError(t, err)
}

func TestGetSize_WithHumanReadable(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := tmpDir + "/test.txt"
	require.NoError(t, os.WriteFile(filepath, make([]byte, 1000000), 0644))
	require.NoError(t, GetPathSize(filepath, Options{HumanReadable: true}))
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

func TestShowHiddenFlag_False(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/normal.txt", make([]byte, 100), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/.hidden.txt", make([]byte, 300), 0644))
	require.NoError(t, GetPathSize(tmpDir, Options{HumanReadable: false, ShowHidden: false}))
}

func TestShowHiddenFlag_True(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/normal.txt", make([]byte, 100), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/.hidden.txt", make([]byte, 300), 0644))
	require.NoError(t, GetPathSize(tmpDir, Options{HumanReadable: false, ShowHidden: true}))
}

func TestShowHiddenFlag_OnlyHiddenFiles(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/.hidden1.txt", make([]byte, 400), 0644))
	require.NoError(t, os.WriteFile(tmpDir+"/.hidden2.txt", make([]byte, 400), 0644))
	require.NoError(t, GetPathSize(tmpDir, Options{HumanReadable: false, ShowHidden: true}))
}

func TestCalcDirSize_OneLevel(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := tmpDir + "/subdir"
	require.NoError(t, os.Mkdir(subDir, 0755))
	require.NoError(t, os.WriteFile(subDir+"/file.txt", make([]byte, 100), 0644))
	size, err := CalcDirSize(tmpDir, Options{Recursive: true})
	require.NoError(t, err)
	require.Equal(t, int64(100), size)
}

func TestCalcDirSize_TwoLevels(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/file.txt", make([]byte, 100), 0644))
	subDir := tmpDir + "/subdir"
	require.NoError(t, os.Mkdir(subDir, 0755))
	require.NoError(t, os.WriteFile(subDir+"/file1.txt", make([]byte, 200), 0644))
	deep := subDir + "/deep"
	require.NoError(t, os.Mkdir(deep, 0755))
	require.NoError(t, os.WriteFile(deep+"/file2.txt", make([]byte, 300), 0644))
	size, err := CalcDirSize(tmpDir, Options{Recursive: true})
	require.NoError(t, err)
	require.Equal(t, int64(600), size)
}

func TestCalcDirSize_HiddenFalse(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/normal.txt", make([]byte, 100), 0644))
	hidden := tmpDir + "/.hidden"
	require.NoError(t, os.Mkdir(hidden, 0755))
	require.NoError(t, os.WriteFile(hidden+"/secret.txt", make([]byte, 500), 0644))
	size, err := CalcDirSize(tmpDir, Options{Recursive: true, ShowHidden: false})
	require.NoError(t, err)
	require.Equal(t, int64(100), size)
}

func TestCalcDirSize_HiddenTrue(t *testing.T) {
	tmpDir := t.TempDir()
	require.NoError(t, os.WriteFile(tmpDir+"/normal.txt", make([]byte, 100), 0644))
	hidden := tmpDir + "/.hidden"
	require.NoError(t, os.Mkdir(hidden, 0755))
	require.NoError(t, os.WriteFile(hidden+"/secret.txt", make([]byte, 500), 0644))
	size, err := CalcDirSize(tmpDir, Options{Recursive: true, ShowHidden: true})
	require.NoError(t, err)
	require.Equal(t, int64(600), size)
}
