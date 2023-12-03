package shortener

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/99pouria/url-shortener/internal/shortener/converter"
)

// Shortener keeps instance of url-shortener
type Shortener struct {
	currentStep atomic.Uint32
	hostname    string
}

// NewShortener creates new 'Shortener' pointer which starts to generate short urls
// by converting given step to a number based 62.
//
// Negative start step replaces with zero and there is no validation for hostname.
// Make sure you are giving appropriate one.
func NewShortener(hostname string, startStep int) *Shortener {
	if startStep < 0 {
		startStep = 0
	}

	s := Shortener{hostname: hostname}
	s.currentStep.Store(uint32(startStep))
	return &s
}

// GenerateNewKey
func (s *Shortener) GenerateNewKey() (string, error) {
	query, err := converter.ItoA(s.currentStep.Add(1))
	if err != nil {
		return "", err
	}
	fmt.Println(query)
	return url.JoinPath(s.hostname, query)
}
