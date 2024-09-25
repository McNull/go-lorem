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

func Pick[T any](r *rand.Rand, arr []T) T {
	l := len(arr)
	idx := r.IntN(l)
	return arr[idx]
}

func PickMultiple[T any](r *rand.Rand, arr []T, n int) []T {
	if n < 1 {
		return []T{}
	}

	x := make([]T, n)
	l := len(arr)

	for i := 0; i < n; i++ {
		x[i] = arr[r.IntN(l)]
	}

	return x
}

func GenerateUniqueRandomInts(r *rand.Rand, max, n int) []int {
	if n > max+1 {
		n = max + 1
	}
	if n < 1 {
		return []int{}
	}

	// Create a slice containing a range of integers from 0 to max
	arr := make([]int, max+1)
	for i := 0; i <= max; i++ {
		arr[i] = i
	}

	// Shuffle the slice using the Fisher-Yates shuffle algorithm

	for i := max; i > 0; i-- {
		j := r.IntN(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	// Return the first n elements of the shuffled slice
	return arr[:n]
}

func PickUnique[T any](r *rand.Rand, arr []T, n int) []T {
	l := len(arr)
	if n > l {
		n = l
	}
	if n < 1 {
		return []T{}
	}

	// Create a slice containing a range of integers from 0 to len(arr) - 1
	indices := GenerateUniqueRandomInts(r, l-1, l)

	x := make([]T, n)

	for i := 0; i < n; i++ {
		x[i] = arr[indices[i]]
	}

	return x
}
