package calc

type BusinessMessage struct {
	Coordinates []Point
	Strategy    int
}

func (r BusinessMessage) Request() BusinessRequest {
	return BusinessRequest(r)
}

func (r BusinessMessage) Response() BusinessResponse {
	return BusinessResponse{}
}

type BusinessRequest struct {
	Coordinates []Point `json:"coordinates"`
	Strategy    int     `json:"strategy"`
}

type BusinessResponse struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Err      string  `json:"err,omitempty"`
}
