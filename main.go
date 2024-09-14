package main

import (
	"fmt"
	"log"
	"os"
	"simalend/nonSharedMap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите путь к файлу в качестве аргумента.")
		return
	}

	filePath := os.Args[1]
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("Файл не найден")
		return
	} else if err != nil {
		fmt.Println(err.Error())
		return
	}

	wordCount, err := nonSharedMap.CountWords(filePath)
	if err != nil {
		log.Fatal(err)
	}
	nonSharedMap.PrintTopWords(os.Stdout, wordCount)

}
