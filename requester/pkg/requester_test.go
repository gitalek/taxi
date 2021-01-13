package requester

//import (
//	a "github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestRequesterService_TripMetrics(t *testing.T) {
//	type args struct {
//		c []Point
//	}
//	tests := map[string]struct{
//		Point Point
//		Lat int
//		Lon int
//	}{
//		"actual price EQUALS minimal one": {Point{1, 2}, 1, 2},
//		"actual price is LESS than minimal one": {3, 2, 150},
//		"actual price is MORE than minimal one": {5, 3, 160},
//	}
//
//	assert := a.New(t)
//	s := new(RequesterService)
//
//	for name, testCase := range tests {
//		t.Run(name, func(t *testing.T) {
//			tm, d, want := testCase.Time, testCase.Dist, testCase.Want
//			got := s.TripMetrics(tm, d)
//			assert.Equal(want, got)
//		})
//	}
//}
