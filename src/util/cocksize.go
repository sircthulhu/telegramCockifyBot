package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateCockSize generates random number with normal distribution N(15, 10*10)
func GenerateCockSize() int {
	size := int(math.Round(rnd.NormFloat64()*15 + 10))
	// size cannot be less or equals to zero
	if size <= 0 {
		size = 1
	}

	return size
}

// FormatCockSizeMessage formats string for cock size
func FormatCockSizeMessage(size int) string {
	str := fmt.Sprintf("My cock size is %d cm.", size)

	if size <= 5 {
		str += " It's cold today 🥶"
	} else if size < 15 { // (5;15)
		str += " Size doesn't matter🫡"
	} else if size > 30 { // (30; infinity)
		str += " Holy shit😱"
	} else if size >= 25 { // [15; 25]
		str += " I'm a giant!😏"
	}

	return str
}
