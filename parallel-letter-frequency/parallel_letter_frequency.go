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

// ConcurrentFrequency counts the frequency of each rune in a given slice of texts and returns this
// data as a FreqMap. Does this work in concurent manner
func ConcurrentFrequency(list []string) FreqMap {
	resCh := make(chan FreqMap, 3)
	for _, s := range list {
		go func(str string) {
			resCh <- Frequency(str)
		}(s)
	}

	res := FreqMap{}

	for range list {
		for letter, freq := range <-resCh {
			res[letter] += freq
		}
	}
	return res
}
