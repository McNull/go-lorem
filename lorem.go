// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import (
	"math/rand/v2"
	"strings"
)

type Lorem struct {
	rnd *rand.Rand
}

func New() *Lorem {
	return &Lorem{rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))}
}

func NewSeeded(seed1 uint64, seed2 uint64) *Lorem {
	return &Lorem{rand.New(rand.NewPCG(seed1, seed2))}
}

// Generate a word in a specfied range of letters.
func (l *Lorem) Word(min, max int) string {
	n := l.intRange(min, max)
	return l.word(n)
}

// Generate a sentence with a specified range of words.
func (l *Lorem) Sentence(min, max int) string {
	n := l.intRange(min, max)
	// grab some words
	ws := []string{}
	maxcommas := 2
	numcomma := 0
	for i := 0; i < n; i++ {
		ws = append(ws, (l.word(l.genWordLen())))

		// maybe insert a comma, if there are currently < 2 commas, and
		// the current word is not the last or first
		if (l.rnd.Int()%n == 0) && numcomma < maxcommas && i < n-1 && i > 2 {
			ws[i-1] += ","
			numcomma += 1
		}

	}

	ws[0] = strings.Title(ws[0])
	sentence := strings.Join(ws, " ") + "."
	return sentence
}

// Generate a paragraph with a specified range of sentenences.
const (
	minwords = 5
	maxwords = 22
)

func (l *Lorem) Paragraph(min, max int) string {
	n := l.intRange(min, max)

	p := []string{}
	for i := 0; i < n; i++ {
		p = append(p, l.Sentence(minwords, maxwords))
	}
	return strings.Join(p, " ")
}

// Generate a random URL
func (l *Lorem) Url() string {
	n := l.intRange(0, 3)

	base := `http://www.` + l.Host()

	switch n {
	case 0:
		break
	case 1:
		base += "/" + l.Word(2, 8)
	case 2:
		base += "/" + l.Word(2, 8) + "/" + l.Word(2, 8) + ".html"
	}
	return base
}

// Host
func (l *Lorem) Host() string {
	n := l.intRange(0, 3)
	tld := ""
	switch n {
	case 0:
		tld = ".com"
	case 1:
		tld = ".net"
	case 2:
		tld = ".org"
	}

	parts := []string{l.Word(2, 8), l.Word(2, 8), tld}
	return strings.Join(parts, ``)
}

// Email
func (l *Lorem) Email() string {
	return l.Word(4, 10) + `@` + l.Host()
}

func (l *Lorem) UUID() string {
	id := RandomUUID(l.rnd)
	return id.String()
}
