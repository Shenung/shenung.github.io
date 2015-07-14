package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func LongestWord(rdr io.Reader) string {
	currentLongestWord := ""
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordOrManyWords := scanner.Text()
		wordOrManyWords = strings.Replace(wordOrManyWords, "?", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "-", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "/", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "*", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, ".", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, ",", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "\"", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "'", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "!", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "?", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "_", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "=", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "+", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "&", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "%", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "^", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "(", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, ")", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "$", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "<", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, ">", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "{", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "}", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "[", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords, "]", " ", -1)
		for _, word := range strings.Fields(wordOrManyWords) {
			if len(word) > len(currentLongestWord) {
				currentLongestWord = word
			}
		}
	}
	return currentLongestWord
}

func main() {
	srcFile, err := os.Open("mobydick.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	longest := LongestWord(srcFile)
	fmt.Println(longest)
}

//
// func WordCount(str string) map[string]int {
//     newStr := strings.ToLower(str)
//     counts := map[string]int{}
//     for _, word := range strings.Fields(newStr) {
//         counts[word]++
//     }
//     return counts
// }
//
// func main() {
//		 fmt.Println("Enter name and book:")
//     var word = os.Args[1]
// srcFile, err := os.Open(os.Args[2])
// if err != nil {
// 		log.Fatalln(err)
// }
// defer srcFile.Close()
//     str := string(srcFile)
//     result := WordCount(str)
//     fmt.Println("Number of", word, ":", result[word])
// }
