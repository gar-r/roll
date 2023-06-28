package dice

// Dice represents a single game dice
type Dice uint

const (
	D4   Dice = 4
	D6   Dice = 6
	D8   Dice = 8
	D10  Dice = 10
	D12  Dice = 12
	D20  Dice = 20
	D100 Dice = 100
)

func SupportedDice() map[int]Dice {
	return map[int]Dice{
		int(D4):   D4,
		int(D6):   D6,
		int(D8):   D8,
		int(D10):  D10,
		int(D12):  D12,
		int(D20):  D20,
		int(D100): D100,
	}
}
