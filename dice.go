// Copyright 2023 by the author. All rights reserved.
// Use of this source code is governed by the Apache
// license 2.0 that can be found in the LICENSE file.
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
	if chance >= 100 {
		return true
	}
	if chance <= 0 {
		return false
	}
	if Roll(100) <= chance {
		return true
	}
	return false
}
