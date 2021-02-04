// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func topWords(path string, K int) []WordCount {
	// Your code here.....
	var e error = nil
	var line []byte
	var words []string

	// Open the file at given path
	file, err := os.Open(path)

	//return error if file cannot be opened
	if err != nil {
		panic("Could not open the file")
	}

	// file will be closed before exiting the function
	defer file.Close()

	reader := bufio.NewReader(file)

	//read the file line by line until EOF encountered
	m := make(map[string]int)
	for e == nil {
		line, _, e = reader.ReadLine()
		words = strings.Fields(string(line))
		for _, word := range words {
			m[word]++
		}
	}
	if e != io.EOF {
		panic("Failed to read the file")
	}

	count := make([]WordCount, 0, len(m))

	for k, v := range m {
		w := WordCount{Word: k, Count: v}
		count = append(count, w)
	}
	sortWordCounts(count)

	//return top k most frequently appearing words
	return count[:K]

}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
