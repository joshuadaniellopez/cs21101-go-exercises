package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Increment(key string, value int) {
	c.mu.Lock()
	c.v[key] += value
	c.mu.Unlock()
}

func (c *SafeCounter) Value() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func main() {
	args := os.Args[1:]
	ch := make(chan map[string]int)
	overallCount := SafeCounter{v: make(map[string]int)}

	for i := 0; i < len(args); i++ {
		go countWords(args[i], ch)
	}
	for i := 0; i < len(args); i++ {
		totalCount := <-ch
		for key, value := range totalCount {
			go overallCount.Increment(key, value)
		}
	}
	time.Sleep(time.Second)

	overallCountMap := overallCount.Value()

	for key, value := range overallCountMap {
		fmt.Printf("%v %v\n", key, value)
	}

}

func countWords(filePath string, ch chan map[string]int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalCount := map[string]int{}
	c := make(chan map[string]int)

	totalLines := 0
	for scanner.Scan() {
		sentence := scanner.Text()
		go countWordsPerLine(sentence, c)
		totalLines += 1
	}

	for i := 0; i < totalLines; i++ {
		count := <-c
		for key, value := range count {
			totalCount[key] += value
		}
	}

	ch <- totalCount
}

func countWordsPerLine(sentence string, c chan map[string]int) {
	sentence = strings.ToLower(sentence)
	f := func(c rune) bool {
		return unicode.IsNumber(c) || unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)
	}
	tokenatedSentence := strings.FieldsFunc(sentence, f)

	count := map[string]int{}
	for i := 0; i < len(tokenatedSentence); i++ {
		count[tokenatedSentence[i]] += 1
	}
	c <- count
}
