package shortid

import "math"

type Random struct {
	previousSeed float64
	seed         float64
}

func newRandom() *Random {
	return &Random{
		seed:         0,
		previousSeed: 0,
	}
}

func (r *Random) setSeed(newSeed float64) {
	r.seed = newSeed
}

func (r *Random) setPrevSeed(seed float64) {
	r.previousSeed = seed
}

func (r *Random) isNewSeed(seed float64) bool {
	if r.previousSeed == seed {
		return false
	}
	return true
}

func (r *Random) random() float64 {
	v := r.seed*9301 + 49297
	f := v - math.Floor(v)
	r.seed = float64(int(v)%233280) + f
	return r.seed / 233280
}
