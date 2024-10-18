package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(list []string) FreqMap {
	freqMap := FreqMap{}
	freqMapStream := make(chan FreqMap, len(list))
	for _, char := range list {
		go func(char string) {
			freqMapStream <- Frequency(char)
		}(char)
	}
	for range list {
		for char, count := range <-freqMapStream {
			freqMap[char] += count
		}
	}
	return freqMap
}
