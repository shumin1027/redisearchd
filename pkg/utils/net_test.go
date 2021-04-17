package utils

import "testing"

func TestResolveDefaultIP(t *testing.T) {
	ip := ResolveDefaultIP()
	println(ip)
}

//func TestDiscoverDefaultIPDefaultIP(t *testing.T) {
//	ip := DiscoverDefaultIP()
//	println(ip)
//}
