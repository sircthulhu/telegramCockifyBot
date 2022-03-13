package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateCockSize generates random number in range [1; 50]
func GenerateCockSize() int {
	// Let x has exponential distribution E(0.4)
	x := rnd.Float64() * 5 // Random number in range [0; 5]
	// then our cock size will be 1 - e ^ (-0.4 * x)
	coefficient := 1 - math.Pow(math.E, -0.4*x)

	// We need to return int as a cock size, so rounding it to the nearest integer
	size := int(math.Round(coefficient * 40))
	if size == 0 {
		size++
	}
	return size
}

// FormatCockSizeMessage formats string for cock size
func FormatCockSizeMessage(size int) string {
	//TODO: emoji
	return fmt.Sprintf("My cock size is %d cm", size)
}
