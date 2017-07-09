package random

import (
    "math/rand"
    "time"
)

// Return: A random range of numbers as a slice within a range
//
// Usage: 
// import "github.com/alok87/goutils/random"
// arr := random.RangeInt(2, 100, 3)
//
// Output: [40, 5, 81]
// Output: [12, 52, 31]
func RangeInt(min int, max int, n int) []int {
     arr := make([]int, n)
     var r int
     for r = 0; r <= n-1; r++ {
         rand.Seed(time.Now().UnixNano())
         arr[r] = rand.Intn(max) + min
     }
     return arr
}
