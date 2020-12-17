// service business logic
package calc

// service as an interface
type Service interface {
	CalculatePrice(int, int) int
	TripMetrics([]Point) (int, int)
}

const (
	taxiService = 50
	minPrice    = 150
	minuteRate  = 10
	kmRate      = 20
)

// implementation of the interface
type CalcService struct{}

// CalculatePrice calculate a price of the trip in rubles (int);
// params: t - number of minutes (int), dist - number of kilometers (int)
func (s *CalcService) CalculatePrice(t int, dist int) int {
	actualPrice := taxiService + t*minuteRate + dist*kmRate
	if minPrice >= actualPrice {
		return minPrice
	}
	return actualPrice
}

// tripMetrics is a temporary stub method until API2 realization
func (*CalcService) TripMetrics(c []Point) (int, int) {
	return int(c[0].Lat), int(c[1].Lat)
}
