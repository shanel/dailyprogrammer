package main
import (
    "bufio"
    "fmt"
    "log"
    "os"  
    "strconv"
    // "strings"
)

func chop(s string) string {
    return s[:len(s)-1]
}

func main() {
    file, err := os.Open("sand.input") // For read access.
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    count := 0
    i := 0
    var grid [][]rune
    for scanner.Scan() {
        text := scanner.Text()
        if count == 0 {
            count, _ = strconv.Atoi(text)
            grid = make([][]rune, count)
            for i := range grid {
                grid[i] = make([]rune, count)
            }
            continue
        }
        for j, val := range text {
            // fmt.Printf("%q", val)
            grid[i][j] = val
        }
        i++
    }
    for _, val := range grid {
        fmt.Printf("%q\n", val)
    }
    // fmt.Printf("%q", grid)
    // quoteme := chop("abcdefg\n")
    // for _, letter := range quoteme {
    //     fmt.Printf("%q\n", letter)
    // }
}