// docdist1.go - initial version of document distance
// Original version by Ronald L. Rivest on February 14, 2007
// Revision by Erik D. Demaine on September 12, 2011
//
// Usage:
//    go run docdist1.go filename1 filename2
// 
// This program computes the "distance" between two text files
// as the angle between their word frequency vectors (in radians).
//
// For each input file, a word-frequency vector is computed as follows:
//    (1) the specified file is read in
//    (2) it is converted into a list of alphanumeric "words"
//        Here a "word" is a sequence of consecutive alphanumeric
//        characters.  Non-alphanumeric characters are treated as blanks.
//        Case is not significant.
//    (3) for each word, its frequency of occurrence is determined
//    (4) the word/frequency lists are sorted into order alphabetically
//
// The "distance" between two vectors is the angle between them.
// If x = (x1, x2, ..., xn) is the first vector (xi = freq of word i)
// and y = (y1, y2, ..., yn) is the second vector,
// then the angle between them is defined as:
//    d(x,y) = arccos(inner_product(x,y) / (norm(x)*norm(y)))
// where:
//    inner_product(x,y) = x1*y1 + x2*y2 + ... xn*yn
//    norm(x) = sqrt(inner_product(x,x))
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"unicode"
)


type NumStringPair struct {
	word string
	occurrence int
}


func isAlNum(r rune) bool {
	fmt.Println(string(r), unicode.IsDigit(r) || unicode.IsLetter(r))
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}


func contains(list []NumStringPair, item string) int {
	for i, el := range list {
		if el.word == item {
			return i
		}
	}

	return -1
}


// ReadFile reads the text file with the given filename;
// return a list of the lines of text in the file.
func ReadFile(filename string) []string {
	
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fileLines := strings.Split(string(fileBytes), "\n")
	return fileLines
}


// GetWordsFromLineList parses the given list L of
// text lines into words.
// Return list of all words found.
func GetWordsFromLineList(list []string) []string {
	wordList := make([]string, 0)

	for _, line := range list {
		wordsInLine := GetWordsFromString(line)
		wordList = append(wordList, wordsInLine...) 
	}

	return wordList
}


// GetWordsFromString returns a list of the words in the given
// input string, converting each word to lower-case.
func GetWordsFromString(line string) []string {
	wordList := make([]string, 0)
	characterList := make([]rune, 0)
	
	for _, s := range line {
		if isAlNum(s) {
			characterList = append(characterList, s)
		} else if len(characterList) > 0 {
			word := strings.ToLower(string(characterList))
			wordList = append(wordList, word)
			characterList = make([]rune, 0)
			fmt.Println(characterList)
		}
	}

	if len(characterList) > 0 {
		word := strings.ToLower(string(characterList))
		wordList = append(wordList, word)
	}

	return wordList
}


// CountFrequency returns a list giving pairs of form: (word,frequency)
func CountFrequency(wordList []string) []NumStringPair {
	var L []NumStringPair

	for _, word := range wordList {
		if v := contains(L, word); v != -1 {
			L[v].occurrence++
		} else {
			L = append(L, NumStringPair{word, 1})
		}
	}
	
	return L
}


// InsertionSort sorts the frequency pairs
func InsertionSort(freqMap []NumStringPair) {
	length := len(freqMap)
	for i := 1; i < length; i++ {
		key := freqMap[i]
		j := i - 1
		
		for j >= 0 && freqMap[j].word > key.word {
			freqMap[j + 1] = freqMap[j]
			j--
		}

		freqMap[j + 1] = key
	}
}


// WordFrequenciesForFile returns alphabetically sorted list
// of (word,frequency) pairs for the given file.
func WordFrequenciesForFile(filename string) []NumStringPair {
	linelist := ReadFile(filename)
	wordList := GetWordsFromLineList(linelist)
	freqMapping := CountFrequency(wordList)
	fmt.Println(freqMapping)
	InsertionSort(freqMapping)
	fmt.Println(freqMapping)
	return freqMapping
}


// VectorAngle a list of (word,freq) pairs, sorted alphabetically.
// Return the angle between these two vectors.
func VectorAngle(L1 []NumStringPair, L2 []NumStringPair) float64 {
	return 0
}


func main() {
	appArgs := os.Args[1:]

	if len(appArgs) != 2 {
		fmt.Println("Usage: go run docdist.go filename_1 filename_2")
	} else {
		filename1 := appArgs[0]
		filename2 := appArgs[1]
		fmt.Printf("%s %s \n", filename1, filename2)

		WordFrequenciesForFile(filename1)
	}
}