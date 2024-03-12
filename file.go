package gostubs

import (
	"os"
	"path/filepath"
)

func ls(path string) ([]string, error) {
	var fileNames = []string{}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range dirEntries {
		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil

}

func lsRecursive(path string) ([]string, error) {
	var fileNames = []string{}
	err := filepath.Walk(path,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return fileNames, nil
}
