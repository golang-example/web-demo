package util

import (
	"os"
	"io/ioutil"
	. "web-demo/log"
)

func ReadFileToString(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		Log.Error(err.Error())
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
