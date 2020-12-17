package calc

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcService_CalculatePrice(t *testing.T) {
	tests := map[string]struct{
		Time int
		Dist int
		Want int
	}{
		"actual price EQUALS minimal one": {4, 3, 150},
		"actual price is LESS than minimal one": {3, 2, 150},
		"actual price is MORE than minimal one": {5, 3, 160},
	}

	assert := a.New(t)
	s := new(CalcService)

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			tm, d, want := testCase.Time, testCase.Dist, testCase.Want
			got := s.CalculatePrice(tm, d)
			assert.Equal(want, got)
		})
	}
}
