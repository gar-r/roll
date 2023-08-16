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
			[]int{4, 4, 3, -4, 5, 6, -3, 7, -2, -2},
			func() int { return rng.Intn(-10, 10) },
		)
	})

	t.Run("lower/upper bound honored (probably)", func(t *testing.T) {
		rng := NewPseudoRandom()
		for i := 0; i < 100000; i++ {
			r := rng.Intn(1, 5)
			assert.GreaterOrEqual(t, r, 1)
			assert.Less(t, r, 5)
		}
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
