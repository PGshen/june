package utils

import (
	"strconv"
	"strings"
)

func Strlist2str(list []string) string {
	var str string
	for e := range list {
		str = str + list[e] + ","
	}
	strings.TrimSuffix(str, ",")
	return str
}

func Intlist2str(list []int32) string {
	var str string
	for e := range list {
		str = str + strconv.Itoa(int(list[e])) + ","
	}
	strings.TrimSuffix(str, ",")
	return str
}

func Intlist2strlist(list []string) []int32 {
	var strlist []int32
	for e := range list {
		temp, _ := strconv.Atoi(list[e])
		strlist = append(strlist, int32(temp))
	}
	return strlist
}
