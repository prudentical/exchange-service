package util

import (
	"path"
	"runtime"
)

func RootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	return dir
}
