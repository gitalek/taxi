package calc
//
//import (
//	"context"
//	a "github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestCalcService_CalculatePrice(t *testing.T) {
//	tests := map[string]struct {
//		Time float64
//		Dist float64
//		Want float64
//	}{
//		"actual price EQUALS minimal one":       {4, 3, 150},
//		"actual price is LESS than minimal one": {3, 2, 150},
//		"actual price is MORE than minimal one": {5, 3, 160},
//	}
//
//	assert := a.New(t)
//	serviceConfig := ServiceConfig{
//		ApiUrl:           "http://localhost:9091/tripmetrics",
//		TaxiServicePrice: 50,
//		MinPrice:         150,
//		MinuteRate:       10,
//		MeterRate:        20,
//	}
//	s := NewCalcService(serviceConfig)
//
//	for name, testCase := range tests {
//		t.Run(name, func(t *testing.T) {
//			tm, d, want := testCase.Time, testCase.Dist, testCase.Want
//			ctx := context.Background()
//			got := s.calculatePrice(ctx, tm, d)
//			assert.Equal(want, got)
//		})
//	}
//}
