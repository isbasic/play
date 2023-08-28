package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"play/common"
	"strings"
)

func GetSevenExecuteFileName() string {
	osName := common.GetOS()
	if osName == "windows" {
		return "7z"
	} else if osName == "darwin" {
		return "7zz"
	} else {
		return "7z"
	}
}

// false 表示路径有重复
func CheckUrlRepeat(fp string) (bool, error) {
	var nameDict map[string]bool
	var e error

	nameList := strings.Split(fp, common.GetSep(common.GetOS()))

	nameDict = make(map[string]bool, len(nameList))

	for _, name := range nameList {
		if _, ok := nameDict[name]; ok {
			e = errors.New(fmt.Sprintf("%s repeat", name))
			return false, e
		} else {
			nameDict[name] = true
		}
	}
	return true, e
}

// 适用于路径中目录名称不重复的情况，如果重复了会产生同名压缩文件的问题
func MakeCompressName(fp string) (string, error) {
	var e error
	var res string
	sep := "_"

	if !common.Exist(fp) {
		e = errors.New("MakeCompressName: File not exist.")
		return res, e
	}

	fileTime, e := common.GetFileTime(fp)
	if e != nil {
		return res, e
	}

	tail := common.FormatTime(fileTime) + ".7z"

	dstName := strings.ReplaceAll(fp, ".", "_")

	filename := fmt.Sprintf("%s%s%s", dstName, sep, tail)
	return filename, nil
}

func CompressFile(fp string) (bool, error) {
	var e error

	src := fp

	if !filepath.IsAbs(src) {
		src, e = filepath.Abs(src)
		if e != nil {
			return false, errors.New(fmt.Sprintf("isabs: %s", e.Error()))
		}
	}

	dst, dstErr := MakeCompressName(src)

	if dstErr != nil {
		fmt.Printf("dstErr: %s", dstErr)
		return false, dstErr
	}

	fmt.Sprintf("CompressFile dst: %s", dst)

	var cmd *exec.Cmd
	osName := common.GetOS()

	// example:"{7z|7zz} a -sdel filename.7z filename"
	if osName == "windows" {
		cmd = exec.Command("cmd", "/C", GetSevenExecuteFileName(), "a", "-sdel", dst, src)
	} else {
		cmd = exec.Command(GetSevenExecuteFileName(), "a", "-sdel", dst, src)
	}

	_, cmdErr := cmd.CombinedOutput() // cmdResult,cmdError

	if cmdErr != nil {
		var e error
		e = errors.New(fmt.Sprintf("cmd error: %s", cmdErr.Error()))

		return false, e
	}

	return true, nil
}
