package secret

const (
	Wink = 1 << iota
	DoubleBlink
	CloseEyes
	Jump
	Reverse
)

func Handshake(code uint) (out []string) {
	if code&Wink != 0 {
		out = append(out, "wink")
	}
	if code&DoubleBlink != 0 {
		out = append(out, "double blink")
	}
	if code&CloseEyes != 0 {
		out = append(out, "close your eyes")
	}
	if code&Jump != 0 {
		out = append(out, "jump")
	}
	if code&Reverse != 0 {
		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
	}
	return
}
