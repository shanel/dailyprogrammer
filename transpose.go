package main

import (
    "bufio"
    // "bytes"
    "fmt"
    "os"
    "strconv"
)

func main() {
    fh, err := os.Open("transpose.txt")
    if err != nil {
        panic("Could not open file!")
    }
    scanner := bufio.NewScanner(fh)
    scanner.Scan()
    count, _ := strconv.Atoi(scanner.Text())
    lines := [][]byte{}
    for scanner.Scan() {
        text := scanner.Text()
        lines = append(lines, []byte(text))
    }
    // // Need to pad
    longest := 0
    for i := range lines {
        if len(lines[i]) > longest {
            longest = len(lines[i])
        }
    }
    out := make([][]byte, longest)
    for i := 0; i < count; i++ {
        for j := 0; j < longest; j++ {
            if j < len(lines[i]) {
                out[j] = append(out[j], lines[i][j])
            } else {
                out[j] = append(out[j], ' ')
            }
        }
    }
    for i := range out {
        fmt.Printf("%q\n", out[i])
    }
}

