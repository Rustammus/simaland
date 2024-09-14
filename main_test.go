package main

import (
	"os"
	"simalend/asyncBlock"
	"simalend/asyncLine"
	"simalend/nonSharedMap"
	"simalend/singleThread"
	"testing"
)

const outFile = "test_out.txt" // "test_out.txt"
const inFile = "big.txt"       //"War_and_Peace.txt" "big.txt"

func BenchmarkSingleThread(b *testing.B) {
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		wordCount, err := singleThread.CountWords(inFile)
		if err != nil {
			b.Error(err)
		}
		singleThread.PrintTopWords(f, wordCount)
	}
}

func BenchmarkAsyncLine(b *testing.B) {
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		wordCount, err := asyncLine.CountWords(inFile)
		if err != nil {
			b.Error(err)
		}
		asyncLine.PrintTopWords(f, wordCount)
	}
}

func BenchmarkAsyncBlock(b *testing.B) {
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		wordCount, err := asyncBlock.CountWords(inFile)
		if err != nil {
			b.Error(err)
		}
		asyncBlock.PrintTopWords(f, wordCount)
	}
}

func BenchmarkNonSharedMap(b *testing.B) {
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		wordCount, err := nonSharedMap.CountWords(inFile)
		if err != nil {
			b.Error(err)
		}
		nonSharedMap.PrintTopWords(f, wordCount)
	}
}
