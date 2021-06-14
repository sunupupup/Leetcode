package main

import "strings"

//击败100 和 8
func replaceSpace1(s string) string {
	//n := len(s)
	ret := ""
	for _, c := range s {
		if c == ' ' {
			ret += "%20"
		} else {
			ret += string(c)
		}
	}
	return ret
}

//直接利用api
//击败 100  和  66
func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
