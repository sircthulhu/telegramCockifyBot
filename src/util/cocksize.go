package util

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateCockSize generates random number with normal distribution N(8, 5)
func GenerateCockSize() int {
	return rnd.Intn(40)
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
