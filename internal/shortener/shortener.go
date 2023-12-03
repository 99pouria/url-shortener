package shortener

import (
	"sync/atomic"

	"github.com/99pouria/url-shortener/internal/shortener/converter"
)

// Shortener keeps instance of url-shortener
type Shortener struct {
	currentStep atomic.Uint32
}

// NewShortener creates new 'Shortener' pointer which starts to generate short urls
// by converting given step to a number based 62.
//
// Negative start step replaces with zero.
func NewShortener(startStep int) *Shortener {
	if startStep < 0 {
		startStep = 0
	}

	s := Shortener{}
	s.currentStep.Store(uint32(startStep))
	return &s
}

// GenerateNewKey retuns unique key
func (s *Shortener) GenerateNewKey() (string, error) { return converter.ItoA(s.currentStep.Add(1)) }
