package types

import (
	"context"
	"net/http"
)

type Point struct {
	Lat float64
	Lon float64
}

type RequestMetrics func(context.Context, []Point, string, string, *http.Client) (float64, float64, error)

type Requester func(context.Context, []Point, *http.Client) (float64, float64, error)

type Strategy func(context.Context, []Point, map[string]Requester, *http.Client) (float64, float64, error)

type MapConfig struct {
	Url, Token string
}

type MapsConfig map[string]MapConfig

