package "main"

import (
    "fmt"
    "rand"
)

func random(min, max int) int { 
    return rand.Intn(max - min) + min 
}
