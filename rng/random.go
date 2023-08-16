package rng

import (
	"math/rand"
	"time"
)

// RandProvider is the default random number generator the api will use.
var RandProvider = NewPseudoRandom()

// Random provides functions to generate random numbers
type Random interface {
	// Intn generates a random integer in the [lower, upper)
	// half-open interval.
	Intn(lower, upper int) int
}

// PseudoRandom implements the Random interface using the
// built-in rand package.
type PseudoRandom struct {
	r *rand.Rand
}

var randProviderFn = newRand

// NewPseudoRandom returns a new Random backed by the
// PseudoRandom implementation.
func NewPseudoRandom() Random {
	return &PseudoRandom{
		r: randProviderFn(),
	}
}

const errInvalidBounds = "invalid bounds"

// Intn generates a random integer in the [lower, upper)
// half-open interval. If upper < lower it panics.
func (p *PseudoRandom) Intn(lower, upper int) int {
	if upper < lower {
		panic(errInvalidBounds)
	}
	return p.r.Intn(upper-lower) + lower
}

func newRand() *rand.Rand {
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src)
}
