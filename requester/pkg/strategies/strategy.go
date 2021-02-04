package strategies

import (
	"context"
	"errors"
	"github.com/gitalek/taxi/requester/pkg/types"
	"sync"
)

func InitStrategies() []types.Strategy {
	// todo: dynamic?
	return []types.Strategy{
		first,
		second,
		third,
		fourth,
		fifth,
	}
}

// ors
func first(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	return maps["ors"](ctx, points)
}

// bing
func second(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	return maps["bing_maps"](ctx, points)
}

// bing otherwise ors
func third(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	t, dist, err := maps["bing_maps"](ctx, points)
	if err != nil {
		return maps["ors"](ctx, points)
	}
	return t, dist, err
}

type res struct {
	duration float64
	distance float64
	err      error
}

// first win
func fourth(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()
	results := make(chan res, len(maps))
	for _, requester := range maps {
		go func(requester types.Requester, ch chan<- res) {
			select {
			case <-ctx.Done():
				return
			default:
				t, dist, err := requester(ctx, points)
				ch <- res{
					duration: t,
					distance: dist,
					err:      err,
				}
			}
		}(requester, results)
	}
	res := <-results
	return res.duration, res.distance, res.err
}

// average
func fifth(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	results := make(chan res, len(maps))
	var wg sync.WaitGroup
	wg.Add(len(maps))
	for _, requester := range maps {
		go func(wg *sync.WaitGroup, requester types.Requester, ch chan<- res) {
			defer wg.Done()
			t, dist, err := requester(ctx, points)
			ch <- res{
				duration: t,
				distance: dist,
				err:      err,
			}
		}(&wg, requester, results)
	}
	go func(wg *sync.WaitGroup, results chan res) {
		defer close(results)
		wg.Wait()
	}(&wg, results)

	var duration, distance float64
	var successCounter int
	for result := range results {
		if result.err != nil {
			continue
		}
		successCounter += 1
		duration += result.duration
		distance += result.distance
	}
	if successCounter == 0 {
		//todo: wrap errors with cycle
		return 0, 0, errors.New("all requests ended with errors")
	}

	count := float64(len(maps))
	return duration / count, distance / count, nil
}
