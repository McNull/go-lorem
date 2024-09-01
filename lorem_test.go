// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLorem(t *testing.T) {

	t.Run("Seeded lorem", func(t *testing.T) {

		lorem1 := NewSeeded(1, 2)
		lorem2 := NewSeeded(1, 2)

		t.Run("Word", func(t *testing.T) {

			w1 := lorem1.Word(1, 10)
			w2 := lorem2.Word(1, 10)

			if w1 != w2 {
				t.Errorf("Expected %s to equal %s", w1, w2)
			}
		})

		t.Run("Sentence", func(t *testing.T) {

			s1 := lorem1.Sentence(1, 10)
			s2 := lorem2.Sentence(1, 10)

			if s1 != s2 {
				t.Errorf("Expected %s to equal %s", s1, s2)
			}
		})

		t.Run("Paragraph", func(t *testing.T) {

			p1 := lorem1.Paragraph(1, 10)
			p2 := lorem2.Paragraph(1, 10)

			if p1 != p2 {
				t.Errorf("Expected %s to equal %s", p1, p2)
			}
		})

	})

	t.Run("Global lorem", func(t *testing.T) {

		t.Run("UUID", func(t *testing.T) {

			uuid := UUID()

			fmt.Println(uuid)

			if len(uuid) != 36 {
				t.Errorf("Expected %s to have length 36", uuid)
			}

			uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)

			if !uuidRegex.MatchString(uuid) {
				t.Errorf("Expected %s to match regex", uuid)
			}

		})
	})
}
