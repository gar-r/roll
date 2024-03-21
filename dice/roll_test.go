package dice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gar-r/roll/rng"
)

func TestDice_Roll(t *testing.T) {
	tests := []struct {
		d            Dice
		lower, upper int
	}{
		{D4, 1, 5},
		{D6, 1, 7},
		{D8, 1, 9},
		{D10, 1, 11},
		{D12, 1, 13},
		{D20, 1, 21},
		{D100, 1, 101},
	}
	mockRand := setupMocks()
	for _, test := range tests {
		t.Run(fmt.Sprintf("D%d", test.d), func(t *testing.T) {
			_ = test.d.Roll()
			mockRand.AssertCalled(t, "Intn", test.lower, test.upper)
		})
	}
}

func TestDice_RollMany(t *testing.T) {
	mockRand := setupMocks()
	_ = D6.RollMany(6)
	mockRand.AssertNumberOfCalls(t, "Intn", 6)
}

func TestDice_RollSpecial(t *testing.T) {
	_ = setupMocks()
	n := D4.RollSpecial(3, 10)
	assert.Equal(t, 13, n)
}

func TestRoll(t *testing.T) {
	_ = setupMocks()
	n, _ := Roll("3d6+2")
	assert.Equal(t, 5, n)
}

func setupMocks() *mockRandomProvider {
	mockRand := &mockRandomProvider{}
	mockRand.On("Intn", mock.Anything, mock.Anything).Return(1)
	rng.RandProvider = mockRand
	return mockRand
}

type mockRandomProvider struct {
	mock.Mock
}

func (m *mockRandomProvider) Intn(lower, upper int) int {
	call := m.Called(lower, upper)
	return call.Int(0)
}
