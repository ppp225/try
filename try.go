package try

import (
	"fmt"
	"math"
	"time"
)

var (
	// MaxRetries is the maximum number of retries before bailing.
	MaxRetries = 5
	// BackoffExponent controls how long backoff should be. Formula is attempt^BackoffExponent
	BackoffExponent = 2.0
	// ErrMaxRetriesReached is returned when retrues were exhausted
	ErrMaxRetriesReached = fmt.Errorf("exceeded retry limit")
)

// Func represents functions that can be retried.
type Func func(attempt int) (err error)

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(fn Func) error {
	var err error
	attempt := 1
	for {
		err = fn(attempt)
		if err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return fmt.Errorf("%w. Last error: %v", ErrMaxRetriesReached, err)
		}
		n := time.Duration(math.Pow(float64(attempt-1), BackoffExponent))
		time.Sleep(n * time.Second)
	}
	return err
}
