package secret

var actions = [4]string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) (out []string) {
	reverseTrigger := false

	for code > 0 {
		switch {
		case code >= 16:
			code -= 16
			reverseTrigger = true
		case code >= 8:
			code -= 8
			out = append(out, "jump")
		case code >= 4:
			code -= 4
			out = append(out, "close your eyes")
		case code >= 2:
			code -= 2
			out = append(out, "double blink")
		case code >= 1:
			code -= 1
			out = append(out, "wink")
		}
	}

	if !reverseTrigger {
		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
	}

	return
}
