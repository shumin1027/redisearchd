package utils

import "time"

//NowDays 给ldap写入数据使用,获取1970-1-1至今的天数
func NowDays(t *time.Time) int64 {
	if t == nil {
		return 0
	}
	return t.Unix() / 3600 / 24
}
