package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"play/cmd"
	"play/common"
)

func main() {
	cmd.InitFlag()
	fp := cmd.FP
	if !filepath.IsAbs(fp) {
		var e error
		fp, e = filepath.Abs(fp)
		if e != nil {
			fmt.Sprintf("main error: %s", e)
			os.Exit(1)
		}
		// fmt.Sprintf("fp: %s", fp)
	}

	ff, err := os.Stat(fp)

	if err != nil {
		fmt.Sprint("%s", err)
		os.Exit(1)
	}

	if ff.IsDir() {
		files := make(chan string)

		go func() {
			common.DirScan(fp, files)
			close(files)
		}()

		for f := range files {
			if !common.Exist(f) {
				fmt.Println("file(f) not found:", f)
			}
			b, e := cmd.CompressFile(f)
			if e != nil {
				fmt.Println("f cmd error:", e)
			}
			if b {
				fmt.Sprintf("%s compress done!", f)
			}
		}
	} else {
		b, e := cmd.CompressFile(fp)
		if e != nil {
			fmt.Println(e)
		}
		if b {
			fmt.Sprintf("%s compress done!", fp)
		}
	}

}

func readDir(fp string) []string {
	var fl []string
	var ffs []fs.FileInfo
	var e error

	ffs, e = ioutil.ReadDir(fp)
	if e != nil {
		fmt.Println("readdir:", e)
		os.Exit(1)
	}
	for _, ff := range ffs {
		fl = append(fl, filepath.Join(fp, common.GetSep(common.GetOS()), ff.Name()))
	}

	return fl
}
