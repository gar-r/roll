package rng

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPseudoRandom(t *testing.T) {
	rng := NewPseudoRandom()
	assert.NotNil(t, rng)
	assert.IsType(t, &PseudoRandom{}, rng)
}

func TestPseudoRandom_Intn(t *testing.T) {

	// mock provider fn with fixed seed
	randProviderFn = func() *rand.Rand {
		return rand.New(rand.NewSource(0))
	}

	t.Run("panics when upper bound smaller than lower bound", func(t *testing.T) {
		rng := NewPseudoRandom()
		assert.PanicsWithValue(t, errInvalidBounds, func() {
			rng.Intn(5, 1)
		})
	})

	t.Run("generate numbers in interval", func(t *testing.T) {
		rng := NewPseudoRandom()
		assertIntSequence(t,
			[]int{2, -1, -3, -6, 4, -9, 9, 4, -7, 5},
			func() int { return rng.Intn(-10, 10) },
		)
	})

}

func assertIntSequence(t *testing.T, expected []int, rngFn func() int) {
	t.Helper()
	actual := make([]int, len(expected))
	for i := 0; i < len(expected); i++ {
		actual[i] = rngFn()
	}
	assert.Equal(t, expected, actual)
}
