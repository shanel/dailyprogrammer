package main

import (
    "fmt"
    "math"
    // "sync/atomic"
)

// func ProcessCoin(coins chan float64,
//                  counter *uint64,
//                  done chan bool) {
//     for {
//         select {
//         case coin := <-coins:
//             two := math.Floor(coin / 2.0)
//             three := math.Floor(coin / 3.0)
//             four := math.Floor(coin / 4.0)
//             if two > 0.0 {
//                 coins <- two
//             } else {
//                 atomic.AddUint64(counter, 1)
//             }
//             if three > 0.0 {
//                 coins <- three
//             } else {
//                 atomic.AddUint64(counter, 1)
//             }
//             if four > 0.0 {
//                 coins <- four
//             } else {
//                 atomic.AddUint64(counter, 1)
//             }
//             if two == 0.0 && three == 0.0 && four == 0.0 {
//                 done <- true
//                 return
//             }
//         default:
//             fmt.Printf("Waiting...\n")
//         }
//     }
// }

// func Coins(coin float64, counter *uint64) {
//     two := math.Floor(coin / 2.0)
//     three := math.Floor(coin / 3.0)
//     four := math.Floor(coin / 4.0)
//     if two > 0.0 {
//         Coins(two, counter)
//     } else {
//         atomic.AddUint64(counter, 1)
//     }
//     if three > 0.0 {
//         Coins(three, counter)
//     } else {
//         atomic.AddUint64(counter, 1)
//     }
//     if four > 0.0 {
//         Coins(four, counter)
//     } else {
//         atomic.AddUint64(counter, 1)
//     }
// }

func MaxCoins(coin float64, max *float64) {
    two := math.Floor(coin / 2.0)
    three := math.Floor(coin / 3.0)
    four := math.Floor(coin / 4.0)
    total := two + three + four
    if total > *max {
        *max = total
    }
    if two > 0.0 {
        MaxCoins(two, max)
    } 
    if three > 0.0 {
        MaxCoins(three, max)
    }
    if four > 0.0 {
        MaxCoins(four, max)
    }
}

func main() {
    // coins := make(chan float64, 1000)
    // var counter uint64 = 0
    // done1 := make(chan bool)
    // done2 := make(chan bool)
    // done3 := make(chan bool)
    // go ProcessCoin(coins, &counter, done1)
    // go ProcessCoin(coins, &counter, done2)
    // go ProcessCoin(coins, &counter, done3)
    // coins <- 1000.0
    // <-done1
    // <-done2
    // <-done3
    // Coins(1000.0, &counter)
    // fmt.Printf("%d\n", counter)
    var max float64 = 0.0
    MaxCoins(1000.0, &max)
    fmt.Printf("%d\n", max)
}
