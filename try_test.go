package try_test

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/ppp225/try"
)

func TestTryExample(t *testing.T) {
	try.BackoffExponent = -10
	SomeFunction := func() (string, error) {
		return "", nil
	}
	var value string
	err := try.Do(func(attempt int) error {
		var err error
		value, err = SomeFunction()
		return err
	})
	_ = value
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func TestTryExamplePanic(t *testing.T) {
	try.BackoffExponent = -10
	SomeFunction := func() (string, error) {
		panic("something went badly wrong")
	}
	var value string
	err := try.Do(func(attempt int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		value, err = SomeFunction()
		return
	})
	_ = value
	if err != nil {
		//log.Fatalln("error:", err)
	}
}

func TestTryDoSuccessful(t *testing.T) {
	try.BackoffExponent = -10
	is := is.New(t)
	callCount := 0
	err := try.Do(func(attempt int) error {
		callCount++
		return nil
	})
	is.NoErr(err)
	is.Equal(callCount, 1)
}

func TestTryDoFailed(t *testing.T) {
	try.BackoffExponent = -10
	is := is.New(t)
	theErr := fmt.Errorf("something went wrong")
	callCount := 0
	err := try.Do(func(attempt int) error {
		callCount++
		return theErr
	})
	is.Equal(errors.Is(err, try.ErrMaxRetriesReached), true)
	is.Equal(strings.Contains(err.Error(), theErr.Error()), true)
	is.Equal(callCount, 5)
}

func TestTryPanics(t *testing.T) {
	try.BackoffExponent = -10
	is := is.New(t)
	theErr := fmt.Errorf("something went wrong")
	callCount := 0
	err := try.Do(func(attempt int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		callCount++
		if attempt > 2 {
			panic("I don't like three")
		}
		err = theErr
		return
	})
	is.Equal(errors.Is(err, try.ErrMaxRetriesReached), true)
	is.Equal(strings.Contains(err.Error(), "panic: I don't like three"), true)
	is.Equal(callCount, 5)
}

func TestRetryLimit(t *testing.T) {
	try.BackoffExponent = -10
	is := is.New(t)
	err := try.Do(func(attempt int) error {
		return fmt.Errorf("nope")
	})
	is.OK(err)
	is.Equal(errors.Is(err, try.ErrMaxRetriesReached), true)
}
