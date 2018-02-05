package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"unicode/utf8"
)

func WriteNewFile(fpath string, fm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fpath), 0755)
	if err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}

	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close()

	err = out.Chmod(fm)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}

	return nil
}

func CompareDirectories(dir1, dir2 string) error {
	return filepath.Walk(dir1, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("could not read entry dir:\n\t %v", err)
		}
		if f.IsDir() {
			return nil
		}
		runes := []rune(path)
		tmpPath := dir2 + string(runes[utf8.RuneCountInString(dir1):utf8.RuneCountInString(path)])

		if f.Size() == 0 {
			tmpFileInfo, err := os.Stat(tmpPath)
			if err != nil {
				return fmt.Errorf("could not read open %s\n\t %v", tmpPath, err)
			}
			if tmpFileInfo.Size() == 0 {
				return nil
			} else {
				return fmt.Errorf("should be zero %s", tmpFileInfo)
			}
		}

		if !CompareFiles(path, tmpPath) {
			return fmt.Errorf("files don't match %s and %s", path, tmpPath)
		}
		return nil
	})
}

func CompareFiles(path1, path2 string) bool {

	f1, err1 := ioutil.ReadFile(path1)

	if err1 != nil {
		return false
	}

	f2, err2 := ioutil.ReadFile(path2)

	if err2 != nil {
		return false
	}

	return bytes.Equal(f1, f2)
}

func MakeTempDir(dirPath string) (string, error) {
	tmp, err := ioutil.TempDir(dirPath, "scratch")
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", fmt.Errorf("%s: making directory:\n\t %v", dirPath, err)
	}
	return tmp, nil
}
