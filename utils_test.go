package lorem

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("PickMultiple", func(t *testing.T) {

		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		t.Run("Should pick n element(s)", func(t *testing.T) {
			rnd := rand.New(rand.NewPCG(1, 2))
			ul := 10

			for i := 0; i < ul; i++ {
				picked := PickMultiple(rnd, arr, i)

				if len(picked) != i {
					t.Errorf("Expected %d element, got %d", i, len(picked))
				}

				for _, p := range picked {
					if !slices.Contains(arr, p) {
						t.Errorf("Expected %d to be in %v", p, arr)
					}
				}
			}
		})

		t.Run("Should pad with duplicates if n > len(arr)", func(t *testing.T) {
			rnd := rand.New(rand.NewPCG(1, 2))
			ul := 10

			for i := 0; i < ul; i++ {
				picked := PickMultiple(rnd, arr, 20)

				if len(picked) != 20 {
					t.Errorf("Expected 20 element, got %d", len(picked))
				}

				for _, p := range picked {
					if !slices.Contains(arr, p) {
						t.Errorf("Expected %d to be in %v", p, arr)
					}
				}
			}
		})
	})

	t.Run("PickUnique", func(t *testing.T) {

		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		t.Run("Should pick n unique element(s)", func(t *testing.T) {
			rnd := rand.New(rand.NewPCG(133, 2))
			ul := 3

			for i := 0; i < ul; i++ {
				picked := PickUnique(rnd, arr, i)

				if len(picked) != i {
					t.Errorf("Expected %d element, got %d", i, len(picked))
				}

				for _, p := range picked {
					if !slices.Contains(arr, p) {
						t.Errorf("Expected %d to be in %v", p, arr)
					}
				}
			}
		})

		t.Run("Should not pick more than len(arr) element(s)", func(t *testing.T) {
			rnd := rand.New(rand.NewPCG(1, 332))
			n := len(arr) + 100

			picked := PickUnique(rnd, arr, n)

			if len(picked) != len(arr) {
				t.Errorf("Expected %d element, got %d", len(arr), len(picked))
			}
		})

		t.Run("Should contain unique elements", func(t *testing.T) {
			rnd := rand.New(rand.NewPCG(333, 2))
			n := len(arr)

			picked := PickUnique(rnd, arr, n)

			fmt.Printf("Picked: %v\n", picked)

			m := map[int]bool{}

			for _, p := range picked {
				if m[p] {
					t.Errorf("Expected %d to be unique", p)
				}
				m[p] = true
			}
		})
	})
}
