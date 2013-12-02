package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
    "strings"
    "strconv"
    "time"
)

func Random(min, max int) int { 
    return rand.Intn(max - min) + min 
}

func DieSplit(s string) (count,kind int) {
    pieces := strings.Split(s, "d")
    count, _ = strconv.Atoi(pieces[0])
    kind, _ = strconv.Atoi(pieces[1])
    return
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = input[:len(input)-1]
    count, kind := DieSplit(input)
    var output []string
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < count; i++ {
        output = append(output, strconv.Itoa(Random(1, kind)))
    }
    fmt.Println(strings.Join(output, " "))
}