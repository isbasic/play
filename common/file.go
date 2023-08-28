package common

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func Exist(fp string) bool {
	_, err := os.Stat(fp)

	return err == nil || os.IsExist(err)
}

func GetDirBase(fp string) (string, error) {
	abs, err := filepath.Abs(fp)
	if err != nil {
		return "", err
	}
	if Exist(abs) {
		dir := filepath.Dir(abs)
		base := filepath.Base(dir)
		return base, err
	} else {
		e := errors.New(fmt.Sprintf("GetDirBase: %s is not exists."))
		return "", e
	}
}

func GetFileInfo(fp string) fs.FileInfo {
	ff, _ := os.Stat(fp)
	return ff
}

func GetFileTime(fp string) (time.Time, error) {
	if !Exist(fp) {
		return time.Now(), errors.New("File not exists, please check %s is valid.")
	}

	ff, err := os.Stat(fp)
	if err != nil {
		return time.Now(), err
	}

	res := ff.ModTime()
	return res, err
}

func GetSep(osName string) string {
	if osName != "" {
		if osName == "windows" {
			return "\\"
		} else {
			return "/"
		}
	} else {
		o := runtime.GOOS
		return GetSep(o)
	}
}

func DirScan(dir string, files chan<- string) {
	for _, entry := range Dirents(dir) {
		name := strings.ToLower(entry.Name())
		if strings.Contains(name, ".ds_store") {
			continue
		} else if strings.Contains(name, "7z") {
			continue
		}

		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			DirScan(subdir, files)
		} else {

			files <- filepath.Join(dir, entry.Name())
		}
	}
}

func Dirents(dir string) []os.FileInfo {
	var file string
	file = dir
	if !filepath.IsAbs(dir) {
		file, _ = filepath.Abs(dir)
	}
	entries, err := ioutil.ReadDir(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "common file Dirents: %v\n", err)
		return nil
	}
	return entries
}
