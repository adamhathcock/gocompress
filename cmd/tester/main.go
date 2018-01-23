package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/adamhathcock/gocompress/pkg/rar"
	zip "github.com/adamhathcock/gocompress/pkg/zip"
	"github.com/adamhathcock/gocompress"
)

func mkdir(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("%s: making directory: %v", dirPath, err)
	}
	return nil
}

func writeNewFile(fpath string, fm os.FileMode) error {
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

func main() {
	tmp, err := ioutil.TempDir("/Users/adam/Desktop/comics", "archiver")
	if err != nil {
		os.Exit(-1)
		return
	}
	// if files come before their containing folders, then we must
	// create their folders before writing the file
	err = mkdir(tmp)
	if err != nil {
		os.Exit(-1)
		return
	}
	//defer os.RemoveAll(tmp)

	ar := rar.RarReader

	fmt.Print(tmp)
	err = ar.OpenPath("/Users/adam/Desktop/Comics/All-Star Batman v1 (2016)/All-Star Batman Vol. 02 - Ends of the Earth (2017) (Digital) (Zone-Empire).cbr")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}
	var entry gocompress.Entry
	for {
		entry, err = ar.ReadEntry()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
			return
		}
		if entry == gocompress.NilEntry {
			break
		}
		if strings.Contains(entry.Name(), "131") {
			path := filepath.Join(tmp, entry.Name())
			writeNewFile(path, 666)
			writer, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
				return
			}
			err = entry.Write(writer)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
				return
			}
			fmt.Println("Wrote " + entry.Name())
		}
	}

	tmp, err = ioutil.TempDir("/Users/adam/Desktop/comics", "archiver")
	if err != nil {
		os.Exit(-1)
		return
	}
	// if files come before their containing folders, then we must
	// create their folders before writing the file
	err = mkdir(tmp)
	if err != nil {
		os.Exit(-1)
		return
	}
	//defer os.RemoveAll(tmp)

	/*rar := archiver.Rar

	fmt.Print(tmp)
	err = rar.Open("/Users/adam/Desktop/Comics/All-Star Batman v1 (2016)/All-Star Batman Vol. 02 - Ends of the Earth (2017) (Digital) (Zone-Empire).cbr", tmp)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}

	tmp, err = ioutil.TempDir("/Users/adam/Downloads", "duet")
	if err != nil {
		os.Exit(-1)
		return
	}
	// if files come before their containing folders, then we must
	// create their folders before writing the file
	err = mkdir(tmp)
	if err != nil {
		os.Exit(-1)
		return
	}
	//defer os.RemoveAll(tmp)*/

	zipReader := zip.ZipReader

	fmt.Print(tmp)
	err = zipReader.OpenPath("/Users/adam/Downloads/Duet-1-7-0-2.zip")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}
	for {
		entry, err = zipReader.ReadEntry()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
			return
		}
		if entry == gocompress.NilEntry {
			break
		}
		/*path := filepath.Join(tmp, entry.Name())
		writeNewFile(path, 666)
		writer, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
			return
		}
		err = entry.Write(writer)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
			return
		}*/
		if !entry.IsDirectory() {
			fmt.Println("Wrote " + entry.Name())
		}
	}
}
