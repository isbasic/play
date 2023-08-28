package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang/glog"
	_ "github.com/google/uuid"
	"github.com/pborman/uuid"
)

func B64Encode(b []byte) string {
	str := base64.StdEncoding.EncodeToString(b)
	return str
}

func B64Decode(in string) ([]byte, error) {
	sDec, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		glog.Errorf("Error decoding string: %s ", err.Error())
	}

	return sDec, err
}

func UnknowTypeAssert(name string, val interface{}) error {
	var e error
	switch t := val.(type) {
	default:
		e = errors.New(fmt.Sprintf("%s need string data, the val type is %T", name, t))
		return e
		break
	}
	return e
}

func GetUUID(hyphen bool) string {
	uuidWithHyphen := uuid.NewRandom()

	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	if hyphen {
		return uuidWithHyphen.String()
	} else {
		return uuid
	}
}

func FloorCount(fp string, sep string) int64 {
	var fileName string
	fileName = fp

	if !filepath.IsAbs(fileName) {
		fileName, _ = filepath.Abs(fileName)
	}

	split := strings.Split(fileName, sep)
	return int64(len(split))
}

func GetOS() string {
	return runtime.GOOS
}
