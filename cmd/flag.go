package cmd

import (
	"flag"
)

var FP string

func InitFlag() {
	flag.StringVar(&FP, "fp", ".", "filepath , a directory full path.")
}
