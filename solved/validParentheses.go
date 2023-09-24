package main

func isValid(s string) bool {
	parantheses := map[string]string{"(": ")", "{": "}", "[": "]"}

	stack := []string{}

	if len(s) == 1 {
		return false
	}

	for i := 0; i < len(s); i++ {

		if _, v := parantheses[string(s[i])]; v == false && len(stack) == 0 {
			return false
		}

		if _, v := parantheses[string(s[i])]; v == true {
			stack = append(stack, string(s[i]))
		}

		if _, v := parantheses[string(s[i])]; v == false && len(stack) != 0 {
			switch {
			case string(s[i]) == ")" && len(stack) != 0 && stack[len(stack)-1] == "(":
				stack = stack[:len(stack)-1]
			case string(s[i]) == "]" && len(stack) != 0 && stack[len(stack)-1] == "[":
				stack = stack[:len(stack)-1]
			case string(s[i]) == "}" && len(stack) != 0 && stack[len(stack)-1] == "{":
				stack = stack[:len(stack)-1]
			default:
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
