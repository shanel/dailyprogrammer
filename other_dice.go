// Not my code

package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
)

func main() {
    numDice, numSides := 0, 0
    if scan, _ := fmt.Sscanf(os.Args[1], "%dd%d", &numDice, &numSides) ; scan == 2 {
        rand.Seed(time.Now().UnixNano())

        diceChan := make(chan int)
        for i := 0 ; i < numDice ; i++ {
            go func() { diceChan <- rand.Intn(numSides) + 1 }()
        }

        for i := 0 ; i < numDice ; i++ {
            fmt.Print(<-diceChan, " ")
        }
        fmt.Print("\n")
    }
}