package calc

type BusinessMessage struct {
	Coordinates []Point
}

func (r BusinessMessage) Request() BusinessRequest {
	return BusinessRequest{r.Coordinates}
}

type BusinessRequest struct {
	Coordinates []Point `json:"coordinates"`
}

func (r BusinessMessage) Response() BusinessResponse {
	return BusinessResponse{}
}

// TripMetrics response
type BusinessResponse struct {
	Distance int    `json:"distance"`
	Duration int    `json:"duration"`
	Err      string `json:"err,omitempty"`
}
