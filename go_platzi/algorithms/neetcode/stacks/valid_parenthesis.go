package main

func isValid(s string) bool {
	orderMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}

	for _, l := range s {
		switch string(l) {
		case ")", "}", "]":
			closer, _ := orderMap[string(l)]
			if len(stack) == 0 {
				return false
			}
			if closer != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, string(l))
		}
	}
	return len(stack) == 0
}
