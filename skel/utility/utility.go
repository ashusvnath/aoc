package utility

import (
	"math/rand"
	"time"
)

var _rng *rand.Rand

func init() {
	_rng = rand.New(rand.NewSource(time.Now().UnixMicro()))
}
