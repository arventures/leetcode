package _20

func IsValid(str string) bool {
	var stack []rune

	bracketMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	if len(str)%2 != 0 {
		return false
	}

	for _, char := range str {
		if closingBracket, ok := bracketMap[char]; ok {
			stack = append(stack, closingBracket)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
