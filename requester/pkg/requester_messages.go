package requester

type BusinessMessage struct {
	Coordinates []Point
}

func (r BusinessMessage) ORSRequest() ORSRequest {
	coordinates := make([][]float64, 0, len(r.Coordinates))
	for _, point := range r.Coordinates {
		coordinates = append(coordinates, []float64{point.Lat, point.Lon})
	}
	return ORSRequest{Coordinates: coordinates, Language: "ru", Units: "m"}
}

type ORSRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
	Language string `json:"language"`
	Units string `json:"units"`
}

func (r BusinessMessage) ORSResponse() ORSResponse {
	return ORSResponse{}
}

type ORSResponse struct {
	Features []struct {
		Properties struct {
			Summary struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
			} `json:"summary"`
		} `json:"properties"`
	} `json:"features"`
}
