package main

import (
	"bufio"
	"fmt"
	"os"
	"rand"
    "strings"
    "strconv"
)

func random(min, max int) int { 
    return rand.Intn(max - min) + min 
}

func main() {
	// An artificial input source.
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	// Count the words.
    // count := 0
    // for scanner.Scan() {
    //  count++
    // }
    // if err := scanner.Err(); err != nil {
    //  fmt.Fprintln(os.Stderr, "reading input:", err)
    // }
    // fmt.Printf("%d\n", count)
    scanner.Scan()
    text := scanner.Text()
    round := 0
    numbers := make([]int, 4)
    var min int
    var max int
    correct := true
    for line := strings.Split(text, " "); line[0] != "q" {
        if round == 0 {
            min = strconv.Atoi(line[0])
            max = strconv.Atoi(line[1])
        }
        if correct {
            rand.Seed(int64(time.Now().Unix()))
            for i := 0; i < 4; i++ {
                numbers[i] = random(min, max)
            }
        }
    }
    count := len(strings.Split(text, " "))
    fmt.Printf("%d\n", count)
}