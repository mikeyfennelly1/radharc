package parse

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
)

func RecursiveSubDirSearch(dirPath string, pattern string) {
	re := regexp.MustCompile(pattern)

	err := filepath.WalkDir(dirPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && re.MatchString(info.Name()) {
			fmt.Println("Matched Directory:", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
