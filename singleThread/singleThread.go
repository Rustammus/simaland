package singleThread

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

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)

		// Words count
		for _, word := range words {
			str := cleanWord(word)
			wordCount[str]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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
	logger.Println("Топ-10 частых слов: singleThread")
	for i := 0; i < 10 && i < len(freqSlice); i++ {
		logger.Printf("%d. %s: %d\n", i+1, freqSlice[i].word, freqSlice[i].count)
	}
}
