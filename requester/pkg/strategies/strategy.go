package strategies

import (
	"context"
	"github.com/gitalek/taxi/requester/pkg/types"
	"sync"
)

func InitStrategies() []types.Strategy {
	// todo: dynamic?
	var strategies []types.Strategy
	strategies = append(strategies, first)
	strategies = append(strategies, second)
	strategies = append(strategies, third)
	strategies = append(strategies, fourth)
	strategies = append(strategies, fifth)
	return strategies
}

func first(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	return maps["ors"](ctx, points)
}

func second(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	return maps["bing_maps"](ctx, points)
}

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

func fourth(ctx context.Context, points []types.Point, maps map[string]types.Requester) (float64, float64, error) {
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan interface{})
	defer func() {
		cancel()
		close(done)
	}()
	results := make(chan res, len(maps))
	for _, requester := range maps {
		go func(done <-chan interface{}, requester types.Requester, ch chan<- res) {
			select {
			case <-done:
				return
			default:
				t, dist, err := requester(ctx, points)
				ch <- res{
					duration: t,
					distance: dist,
					err:      err,
				}
			}
		}(done, requester, results)
	}
	res := <-results
	return res.duration, res.distance, res.err
}

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
	for result := range results {
		if result.err != nil {
			return 0, 0, result.err
		}
		duration += result.duration
		distance += result.distance
	}
	count := float64(len(maps))
	return duration / count, distance / count, nil
}
