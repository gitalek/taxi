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

func (r BusinessMessage) ORSResponse() ORSResponse {
	return ORSResponse{}
}

type ORSRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
	Language string `json:"language"`
	Units string `json:"units"`
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

type BMResponse struct {
	//AuthenticationResultCode string `json:"authenticationResultCode"`
	//BrandLogoURI             string `json:"brandLogoUri"`
	//Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		//EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			//Type         string    `json:"__type"`
			//Bbox         []float64 `json:"bbox"`
			//ID           string    `json:"id"`
			//DistanceUnit string    `json:"distanceUnit"`
			//DurationUnit string    `json:"durationUnit"`
			//RouteLegs    []struct {
			//	ActualEnd struct {
			//		Type        string    `json:"type"`
			//		Coordinates []float64 `json:"coordinates"`
			//	} `json:"actualEnd"`
			//	ActualStart struct {
			//		Type        string    `json:"type"`
			//		Coordinates []float64 `json:"coordinates"`
			//	} `json:"actualStart"`
			//	AlternateVias  []interface{} `json:"alternateVias"`
			//	Cost           int           `json:"cost"`
			//	Description    string        `json:"description"`
			//	ItineraryItems []struct {
			//		CompassDirection string `json:"compassDirection"`
			//		Details          []struct {
			//			CompassDegrees              int      `json:"compassDegrees"`
			//			EndPathIndices              []int    `json:"endPathIndices"`
			//			LocationCodes               []string `json:"locationCodes"`
			//			ManeuverType                string   `json:"maneuverType"`
			//			Mode                        string   `json:"mode"`
			//			Names                       []string `json:"names"`
			//			RoadShieldRequestParameters struct {
			//				Bucket  int `json:"bucket"`
			//				Shields []struct {
			//					Labels         []string `json:"labels"`
			//					RoadShieldType int      `json:"roadShieldType"`
			//				} `json:"shields"`
			//			} `json:"roadShieldRequestParameters"`
			//			RoadType         string `json:"roadType"`
			//			StartPathIndices []int  `json:"startPathIndices"`
			//		} `json:"details"`
			//		Exit        string `json:"exit"`
			//		IconType    string `json:"iconType"`
			//		Instruction struct {
			//			FormattedText interface{} `json:"formattedText"`
			//			ManeuverType  string      `json:"maneuverType"`
			//			Text          string      `json:"text"`
			//		} `json:"instruction"`
			//		IsRealTimeTransit bool `json:"isRealTimeTransit"`
			//		ManeuverPoint     struct {
			//			Type        string    `json:"type"`
			//			Coordinates []float64 `json:"coordinates"`
			//		} `json:"maneuverPoint"`
			//		RealTimeTransitDelay int     `json:"realTimeTransitDelay"`
			//		SideOfStreet         string  `json:"sideOfStreet"`
			//		TollZone             string  `json:"tollZone"`
			//		TowardsRoadName      string  `json:"towardsRoadName,omitempty"`
			//		TransitTerminus      string  `json:"transitTerminus"`
			//		TravelDistance       float64 `json:"travelDistance"`
			//		TravelDuration       int     `json:"travelDuration"`
			//		TravelMode           string  `json:"travelMode"`
			//		Hints                []struct {
			//			HintType string `json:"hintType"`
			//			Text     string `json:"text"`
			//		} `json:"hints,omitempty"`
			//		Signs    []string `json:"signs,omitempty"`
			//		Warnings []struct {
			//			Origin      string `json:"origin"`
			//			Severity    string `json:"severity"`
			//			Text        string `json:"text"`
			//			To          string `json:"to"`
			//			WarningType string `json:"warningType"`
			//		} `json:"warnings,omitempty"`
			//	} `json:"itineraryItems"`
			//	RouteRegion  string `json:"routeRegion"`
			//	RouteSubLegs []struct {
			//		EndWaypoint struct {
			//			Type               string    `json:"type"`
			//			Coordinates        []float64 `json:"coordinates"`
			//			Description        string    `json:"description"`
			//			IsVia              bool      `json:"isVia"`
			//			LocationIdentifier string    `json:"locationIdentifier"`
			//			RoutePathIndex     int       `json:"routePathIndex"`
			//		} `json:"endWaypoint"`
			//		StartWaypoint struct {
			//			Type               string    `json:"type"`
			//			Coordinates        []float64 `json:"coordinates"`
			//			Description        string    `json:"description"`
			//			IsVia              bool      `json:"isVia"`
			//			LocationIdentifier string    `json:"locationIdentifier"`
			//			RoutePathIndex     int       `json:"routePathIndex"`
			//		} `json:"startWaypoint"`
			//		TravelDistance float64 `json:"travelDistance"`
			//		TravelDuration int     `json:"travelDuration"`
			//	} `json:"routeSubLegs"`
			//	TravelDistance float64 `json:"travelDistance"`
			//	TravelDuration int     `json:"travelDuration"`
			//	TravelMode     string  `json:"travelMode"`
			//} `json:"routeLegs"`
			//TrafficCongestion     string  `json:"trafficCongestion"`
			//TrafficDataUsed       string  `json:"trafficDataUsed"`
			TravelDistance        float64 `json:"travelDistance"`
			TravelDuration        int     `json:"travelDuration"`
			//TravelDurationTraffic int     `json:"travelDurationTraffic"`
			//TravelMode            string  `json:"travelMode"`
		} `json:"resources"`
	} `json:"resourceSets"`
	//StatusCode        int    `json:"statusCode"`
	//StatusDescription string `json:"statusDescription"`
	//TraceID           string `json:"traceId"`
}
