# dice
Roll a dice.

## Functions

### func Roll
```go
func Roll(sides int) int
```
*Roll* rolls a die with a given number of *sides* and returns a pseudo-random number in the range [1-sides]. *Roll* is a simple wrapper function for golang's *math/rand.Intn* function which will panic if *sides* is smaller or equal to 0.

#### Example
```go
package main

import(
    "fmt"
    
    "github.com/aschulth/dice"
)

func main() {
    fmt.Printf("You rolled a D6 and got a %d.\n", dice.Roll(6))
}
```

### func RollChance
```go
func RollChance(chance int) bool
```
*RollChance* rolls a hundred sided die and returns *true* if the result is smaller or equal to *chance*. Otherwise it returns *false*. The value for *chance* may be > 100, in which case *RollChance* returns *true*, or <= 0, in which case *RollChance* returns *false*.

#### Example
```go
package main

import(
    "fmt"
    
    "github.com/aschulth/dice"
)

func main() {
    fmt.Println("# You have an 80% chance to hit the monster.")
    
    if dice.RollChance(80) {
        fmt.Printf("You hit the monster! Now it's angry :S.\n")
    } else {
        fmt.Printf("You missed the monster! Oh oh...\n")
    }
}
```

## Constants
None.

## Variables
None.
