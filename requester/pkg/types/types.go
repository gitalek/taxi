package types

import (
	"context"
)

type Point struct {
	Lat float64
	Lon float64
}

type RequestMetrics func(context.Context, []Point, string, string) (float64, float64, error)

type Requester func(context.Context, []Point) (float64, float64, error)

type Strategy func(context.Context, []Point, map[string]Requester) (float64, float64, error)
