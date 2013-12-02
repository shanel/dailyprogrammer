package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strings"
    "sync"
    "sync/atomic"
)

// might need to lock this somehow - might be race conditions if diff go routines are updating same count
func CountLettersAndSymbols(word string, letters map[string]int, 
                            symbols map[string]int, l_mu *sync.Mutex,
                            s_mu *sync.Mutex, lc *uint64, sc *uint64) {
    for _, r := range word {
        s := strings.ToLower(string(r))
        if m, _ := regexp.MatchString("[a-z]", s); m {
            l_mu.Lock()
            l_mu.Unlock()
            atomic.AddUint64(lc, 1)
        } else {
            s_mu.Lock()
            _, prs := symbols[s]
            if prs {
                symbols[s] = symbols[s] + 1
            } else {
                symbols[s] = 1
            }
            s_mu.Unlock()
            atomic.AddUint64(sc, 1)
        }
    }
}

func CountWords(word string, words map[string]int, w_mu *sync.Mutex) {
    re := regexp.MustCompile("[a-z]+")
    the_word := re.FindString(word)
    w_mu.Lock()
    _, prs := words[the_word]
    if prs {
        words[the_word] = words[the_word] + 1
    } else {
        words[the_word] = 1
    }
    w_mu.Unlock()
}

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func GetTopThree(the_map map[string]int) (one,two,three string) {
    pl := sortedKeys(the_map)
    one = pl[0]
    two = pl[1]
    three = pl[2]
    return
}

func DoTheWork(wordchan chan string, wc *uint64, l_mu *sync.Mutex, s_mu *sync.Mutex,
               w_mu *sync.Mutex, l_map map[string]int,
               s_map map[string]int, w_map map[string]int, done chan bool, lc *uint64, sc *uint64) {
    for {
        w, more := <-wordchan
        if more {
            CountLettersAndSymbols(w, l_map, s_map, l_mu, s_mu, lc, sc)
            atomic.AddUint64(wc, 1)
            CountWords(w, w_map, w_mu)
        } else {
            done <- true
            return
        }
    }
}

func main() {
    var letter_mutex = &sync.Mutex{}
    var symbol_mutex = &sync.Mutex{}
    var word_mutex = &sync.Mutex{}
    var letter_count uint64 = 0
    var symbol_count uint64 = 0
    var word_count uint64 = 0
    symbols := map[string]int{}
    letters := map[string]int{"a": 0,
                              "b": 0,
                              "c": 0,
                              "d": 0,
                              "e": 0,
                              "f": 0,
                              "g": 0,
                              "h": 0,
                              "i": 0,
                              "j": 0,
                              "k": 0,
                              "l": 0,
                              "m": 0,
                              "n": 0,
                              "o": 0,
                              "p": 0,
                              "q": 0,
                              "r": 0,
                              "s": 0,
                              "t": 0,
                              "u": 0,
                              "v": 0,
                              "w": 0,
                              "x": 0,
                              "y": 0,
                              "z": 0}
    words := map[string]int{}
    fh, err := os.Open("lorem.txt")
    if err != nil {
        panic("Couldn't open file!")
    }
    wordchan := make(chan string)
    scanner := bufio.NewScanner(fh)
    scanner.Split(bufio.ScanWords)
    
    done1 := make(chan bool)
    done2 := make(chan bool)
    done3 := make(chan bool)
    
    go DoTheWork(wordchan, &word_count, letter_mutex, symbol_mutex,
                 word_mutex, letters, symbols, words, done1, &letter_count, &symbol_count)
    go DoTheWork(wordchan, &word_count, letter_mutex, symbol_mutex,
                 word_mutex, letters, symbols, words, done2, &letter_count, &symbol_count)
    go DoTheWork(wordchan, &word_count, letter_mutex, symbol_mutex,
                 word_mutex, letters, symbols, words, done3, &letter_count, &symbol_count)
    
    for scanner.Scan() {
        wordchan <- strings.ToLower(scanner.Text())
    }
    close(wordchan)
    <-done1
    <-done2
    <-done3
    w1, w2, w3 := GetTopThree(words)
    l1, l2, l3 := GetTopThree(letters)
    fmt.Printf("%d words\n", word_count)
    fmt.Printf("%d letters\n", letter_count)
    fmt.Printf("%d symbols\n", symbol_count)
    fmt.Printf("Top three most common words: %s, %s, %s\n", w1, w2, w3)
    fmt.Printf("Top three most common letters: %s, %s, %s\n", l1, l2, l3)
}

