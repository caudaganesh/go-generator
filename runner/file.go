package runner

import (
	"os"
	"path/filepath"
)

const fileFullPath = "tmp_gogen/main.go"

func GetFullPath() string {
	return fileFullPath
}

func getDir() string {
	return filepath.Dir(GetFullPath())
}

func CreateFile() (*os.File, error) {
	err := os.Mkdir(getDir(), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(GetFullPath())
}

func DeleteFile() error {
	err := os.Remove(GetFullPath())
	if err != nil {
		return err
	}
	return os.Remove(getDir())
}
