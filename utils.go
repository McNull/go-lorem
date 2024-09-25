package lorem

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

// Generate a natural word len.
func (l *Lorem) genWordLen() int {
	f := l.Randomizer.Float32() * 100
	// a table of word lengths and their frequencies.
	switch {
	case f < 1.939:
		return 1
	case f < 19.01:
		return 2
	case f < 38.00:
		return 3
	case f < 50.41:
		return 4
	case f < 61.00:
		return 5
	case f < 70.09:
		return 6
	case f < 78.97:
		return 7
	case f < 85.65:
		return 8
	case f < 90.87:
		return 9
	case f < 95.05:
		return 10
	case f < 97.27:
		return 11
	case f < 98.67:
		return 12
	case f < 100.0:
		return 13
	}
	return 2 // shouldn't get here
}

func (l *Lorem) intRange(min, max int) int {
	if min == max {
		return l.intRange(min, min+1)
	}
	if min > max {
		return l.intRange(max, min)
	}
	n := l.Randomizer.Int() % (max - min)
	return n + min
}

func (l *Lorem) word(wordLen int) string {
	if wordLen < 1 {
		wordLen = 1
	}
	if wordLen > 13 {
		wordLen = 13
	}

	n := l.Randomizer.Int() % len(wordlist)
	for {
		if n >= len(wordlist)-1 {
			n = 0
		}
		if len(wordlist[n]) == wordLen {
			return wordlist[n]
		}
		n++
	}
}

func RandomUUID(rnd *rand.Rand) uuid.UUID {
	bb := make([]byte, 16)

	for i := 0; i < 16; i++ {
		bb[i] = byte(rnd.IntN(256))
	}

	var u uuid.UUID
	copy(u[:], bb)

	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80

	return u
}
