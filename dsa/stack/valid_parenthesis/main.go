package main

func pop(s []byte) []byte {
	return s[:len(s)-1]

}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var charSlice []byte = make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '}' || s[i] == ')' || s[i] == ']' {
			if len(charSlice) == 0 {
				return false
			}
			switch {
			case (charSlice[len(charSlice)-1] == '{' && s[i] == '}'):
				charSlice = pop(charSlice)
			case (charSlice[len(charSlice)-1] == '[' && s[i] == ']'):
				charSlice = pop(charSlice)
			case (charSlice[len(charSlice)-1] == '(' && s[i] == ')'):
				charSlice = pop(charSlice)
			default:
				return false
			}
		} else {
			charSlice = append(charSlice, s[i])
		}
	}
	if len(charSlice) == 0 {
		return true
	} else {
		return false
	}
}
