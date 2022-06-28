package main

import (
	"errors"
	"github.com/cenkalti/backoff"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func main() {
	TestRetry()
}

func TestRetry() {
	const successOn = 10000
	var i = 0

	// This function is successful on "successOn" calls.
	f := func() error {
		i++
		log.Printf("function is called %d. time\n", i)

		if i == successOn {
			log.Println("OK")
			return nil
		}

		log.Println("error")
		return errors.New("error")
	}
	b := &backoff.ExponentialBackOff{
		InitialInterval:     backoff.DefaultInitialInterval,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         backoff.DefaultMaxInterval,
		MaxElapsedTime:      10 * 364 * 24 * time.Hour,
		Clock:               backoff.SystemClock,
	}
	b.Reset()

	err := backoff.Retry(f, b)
	if err != nil {
		logrus.Errorf("unexpected error: %s", err.Error())
	}
	if i != successOn {
		logrus.Errorf("invalid number of retries: %d", i)
	}
}
