package dice

import (
	"math/rand"
	"time"
)

func Roll(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func RollChance(chance int) bool {
    if chance >= 100 { return true }
    if chance <= 0 { return false }
    if Roll(100) <= chance { return true }
    return false
}
