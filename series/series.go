package series

func All(n int, s string) (out []string) {
	for i := 0; i < len(s)-n+1; i++ {
		out = append(out, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	if first, ok := First(n, s); !ok {
		return ""
	} else {
		return first
	}
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}

// func AskTooMuch
