package dice

import (
	"math/rand"
	"time"
)

func Roll(s int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(s) + 1
}

func RollPerc(percentage int) bool {
    if percentage >= 100 { return true }
    if percentage <= 0 { return false }
    if Roll(100) <= percentage { return true }
    return false
}
