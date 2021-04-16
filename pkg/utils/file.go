package utils

import (
	"io"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDir(path string, perm ...os.FileMode) error {
	if len(perm) == 0 {
		return os.MkdirAll(path, os.ModePerm)
	} else {
		return os.MkdirAll(path, perm[0])
	}
}

func MakeFile(path string) (*os.File, error) {
	i := strings.LastIndex(path, "/")
	if i > 0 {
		subdir := path[0 : i+1]
		exists, _ := PathExists(subdir)
		if !exists {
			_ = MakeDir(subdir)
		}
	}
	return os.Create(path)
}

func CopyFile(srcfile, dstfile string, perm ...os.FileMode) (written int64, err error) {
	src, err := os.Open(srcfile)
	if err != nil {
		return
	}
	defer src.Close()
	var dst *os.File
	if len(perm) == 0 {
		dst, err = os.OpenFile(dstfile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	} else {
		dst, err = os.OpenFile(dstfile, os.O_WRONLY|os.O_CREATE, perm[0])
	}
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
