package httpclient

import (
	"math"
	"math/rand"
	"time"

	"github.com/pubgo/golug/golug_entry/http_entry/httpclient/httpclient"
)

// Backoff interface defines contract for backoff strategies
type Backoff = httpclient.Backoff

type constantBackoff struct {
	backoffInterval       int64
	maximumJitterInterval int64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewConstantBackoff returns an instance of ConstantBackoff
func NewConstantBackoff(backoffInterval, maximumJitterInterval time.Duration) Backoff {
	return &constantBackoff{
		backoffInterval:       int64(backoffInterval / time.Millisecond),
		maximumJitterInterval: int64(maximumJitterInterval / time.Millisecond),
	}
}

// Next returns next time for retrying operation with constant strategy
func (cb *constantBackoff) Next(retry int) time.Duration {
	if retry <= 0 {
		return time.Duration(cb.backoffInterval) * time.Millisecond
	}

	return (time.Duration(cb.backoffInterval) * time.Millisecond) + (time.Duration(rand.Int63n(cb.maximumJitterInterval)) * time.Millisecond)
}

type exponentialBackoff struct {
	exponentFactor        float64
	initialTimeout        float64
	maxTimeout            float64
	maximumJitterInterval int64
}

// NewExponentialBackoff returns an instance of ExponentialBackoff
func NewExponentialBackoff(initialTimeout, maxTimeout time.Duration, exponentFactor float64, maximumJitterInterval time.Duration) Backoff {
	return &exponentialBackoff{
		exponentFactor:        exponentFactor,
		initialTimeout:        float64(initialTimeout / time.Millisecond),
		maxTimeout:            float64(maxTimeout / time.Millisecond),
		maximumJitterInterval: int64(maximumJitterInterval / time.Millisecond),
	}
}

// Next returns next time for retrying operation with exponential strategy
func (eb *exponentialBackoff) Next(retry int) time.Duration {
	if retry <= 0 {
		return time.Duration(eb.initialTimeout) * time.Millisecond
	}

	return time.Duration(math.Min(eb.initialTimeout+math.Pow(eb.exponentFactor, float64(retry)), eb.maxTimeout)+float64(rand.Int63n(eb.maximumJitterInterval))) * time.Millisecond
}
