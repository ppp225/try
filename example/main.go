package main

import (
	"fmt"
	"log"

	"github.com/ppp225/try"
)

func succeedOnAttempt(a, n int) error {
	if a == n {
		return nil
	}
	return fmt.Errorf("some error occurred")
}

func main() {
	log.Printf("main() started")
	err := try.Do(func(attempt int) (err error) {
		err = succeedOnAttempt(attempt, 5)
		if err != nil {
			log.Printf("main() failed %d times on succeedOnAttempt(attempt, 5) with err: %v", attempt, err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("main() failed with err: %v", err)
	}

	log.Printf("main() finished")
}
