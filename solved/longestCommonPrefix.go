package main

import (
	//"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	l := len(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]

}
