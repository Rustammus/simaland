package asyncLine

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

func CountWords(filePath string) (map[string]int, error) {

	// Init reader func
	lineReader := func(out chan<- string, filePath string) {
		defer close(out)

		file, err := os.Open(filePath)
		defer file.Close()
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- strings.ToLower(scanner.Text())
		}
	}

	lineChan := make(chan string, 100)
	go lineReader(lineChan, filePath)

	wordCount := make(map[string]int)

	for line := range lineChan {
		words := strings.Fields(line)
		for _, word := range words {
			wordCount[cleanWord(word)]++
		}
	}

	return wordCount, nil
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
	logger.Println("Топ-10 частых слов: asyncLine")
	for i := 0; i < 10 && i < len(freqSlice); i++ {
		logger.Printf("%d. %s: %d\n", i+1, freqSlice[i].word, freqSlice[i].count)
	}
}
