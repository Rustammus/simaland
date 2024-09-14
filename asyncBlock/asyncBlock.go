package asyncBlock

import (
	"io"
	"log"
	"os"
	"runtime"
	"slices"
	"strings"
	"unicode"
)

func CountWords(filePath string) (map[string]int, error) {

	// Init reader func
	lineReader := func(out chan<- string, filePath string) {
		defer close(out)
		buffer := make([]byte, 4096)
		file, err := os.Open(filePath)
		defer file.Close()
		if err != nil {
			return
		}

		lastTrimmed := make([]byte, 0)
		trimmed := make([]byte, 0)
		str := make([]byte, 0)
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				out <- string(buffer[:n])
				break
			} else if err == nil {
				trimmed, str = trimToWord(buffer[:n])
				out <- string(lastTrimmed) + string(str)
				lastTrimmed = trimmed
			}
			runtime.Gosched()
		}
	}

	lineChan := make(chan string, 12)
	go lineReader(lineChan, filePath)

	wordCount := make(map[string]int)

	for line := range lineChan {
		line = strings.ToLower(line)
		words := strings.Fields(line)
		for _, word := range words {
			cWord := cleanWord(word)
			wordCount[cWord]++
		}
	}

	return wordCount, nil
}

func trimToWord(b []byte) ([]byte, []byte) {
	trimmed := make([]byte, 0)
	for i := len(b) - 1; i >= 0; i-- {
		if unicode.IsSpace(rune(b[i])) {
			slices.Reverse(trimmed)
			return trimmed, b[:i+1]
		} else {
			trimmed = append(trimmed, b[i])
		}
	}
	slices.Reverse(trimmed)
	return trimmed, []byte{}
}

// Return only IsLetter or IsDigit chars
func cleanWord(word string) string {
	var result strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func PrintTopWords(writer io.Writer, wordCount map[string]int) {
	logger := log.New(writer, "", 0)

	type wordFrequency struct {
		word  string
		count int
	}

	// Copy from one map to wordFrequency slice
	freqSlice := make([]wordFrequency, 0, len(wordCount))
	for word, count := range wordCount {
		freqSlice = append(freqSlice, wordFrequency{word, count})
	}

	// Sorting
	slices.SortFunc(freqSlice, func(i, j wordFrequency) int {
		return j.count - i.count
	})

	// Output
	logger.Println("Топ-10 самых частых слов: asyncBlock")
	for i := 0; i < 10 && i < len(freqSlice); i++ {
		logger.Printf("%d. %s: %d\n", i+1, freqSlice[i].word, freqSlice[i].count)
	}
}
