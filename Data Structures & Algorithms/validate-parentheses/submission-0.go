func isValid(s string) bool {
    stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]
 
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]

		if char == ')' && top != '(' {
			return false
		}
		if char == '}' && top != '{' {
			return false
		}
		if char == ']' && top != '[' {
			return false
		}

		stack = stack[:len(stack) - 1]
		}
		
	}
	return len(stack) == 0
}
