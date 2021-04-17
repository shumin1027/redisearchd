package utils

import (
	"io"
	"os"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 判断所给路径是否为文件夹
func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

// 判断所给路径是否为文件
func IsFile(path string) (bool, error) {
	isdir, err := IsDir(path)
	return !isdir, err
}

// 创建文件夹
func MakeDir(path string, perm ...os.FileMode) error {
	if len(perm) == 0 {
		return os.MkdirAll(path, os.ModePerm)
	} else {
		return os.MkdirAll(path, perm[0])
	}
}

// 创建文件
func MakeFile(path string) (*os.File, error) {
	i := strings.LastIndex(path, "/")
	if i > 0 {
		subdir := path[0 : i+1]
		exists, _ := Exists(subdir)
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
