package armDice

import (
	"math/rand"
	"time"
)

func Roll() int {
	return RollCustom(1, 6)
}

func RollCustom(min int, max int) int {
	timeSource := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(timeSource)
	return seed.Intn(max-min+1) + min
}
