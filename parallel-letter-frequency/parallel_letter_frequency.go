package letter

import "sync"

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
func ConcurrentFrequency(text []string) FreqMap {
	m := FreqMap{}
	c := make(chan rune, 1)
	var wg sync.WaitGroup

	// read
	go func() {
		for {
			r, ok := <-c
			if ok {
				m[r]++
			} else {
				return
			}
		}
	}()

	// write
	f := func(s string, c chan rune, wg *sync.WaitGroup) {
		for _, r := range s {
			c <- r
		}
		wg.Done()
	}

	// split job in goroutines
	for _, s := range text {
		wg.Add(1)
		go f(s, c, &wg)
		wg.Wait()
	}
	return m
}
