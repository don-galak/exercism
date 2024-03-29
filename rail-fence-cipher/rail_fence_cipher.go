package railfence

import "bytes"

func Encode(message string, rails int) string {
	var encoded bytes.Buffer
	arr := make([][]rune, rails)
	for i := 0; i < rails; i++ {
		arr[i] = make([]rune, len(message))
	}

	add := true
	j := 0
	for i, char := range message {
		if j == rails-1 {
			add = false
		} else if j == 0 {
			add = true
		}
		arr[j][i] = char
		if add {
			j++
		} else {
			j--
		}
	}

	for i := range arr {
		for _, char := range arr[i] {
			if char != 0 {
				encoded.WriteRune(char)
			}
		}
	}

	return encoded.String()
}

func Decode(message string, rails int) string {
	var encoded bytes.Buffer
	msgCopy := message
	arr := make([][]rune, rails)
	for i := 0; i < rails; i++ {
		arr[i] = make([]rune, len(message))
	}

	add := true
	j := 0
	for i := range message {
		if j == rails-1 {
			add = false
		} else if j == 0 {
			add = true
		}
		arr[j][i] = 1
		if add {
			j++
		} else {
			j--
		}
	}

	for i := range arr {
		for j, char := range arr[i] {
			if char == 1 {
				arr[i][j] = rune(msgCopy[0])
				msgCopy = msgCopy[1:]
			}
		}
	}

	for i := range message {
		for j := range arr {
			if arr[j][i] != 0 {
				encoded.WriteRune(arr[j][i])
			}
		}
	}

	return encoded.String()
}
