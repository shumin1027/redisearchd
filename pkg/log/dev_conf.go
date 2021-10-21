//go:build !release
// +build !release

package log

var (
	Conf = DevLogConf()
)
