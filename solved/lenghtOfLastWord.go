package main

func lengthOfLastWord(s string) int {

	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && n != 0 {
			return n
		} else if s[i] != ' ' {
			n += 1
		}
	}
	return n
}
