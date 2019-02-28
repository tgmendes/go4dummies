package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := uselessFunc()
	if err != nil {
		log.Fatal("an error: ", err.Error())
	}
}

// START OMIT
type RandomError struct {
	Msg string
}

func (re RandomError) Error() string {
	return fmt.Sprintf("this was a random, %s", re.Msg)
}

// END OMIT

func uselessFunc() error {
	if 0 > 1 {
		return errors.New("this doesn't look good")
	}

	if 2 > 1 {
		return &RandomError{
			Msg: "dumb error",
		}
	}

	if 3 > 5 {
		return fmt.Errorf("expected %i to be greater than %i", 5, 3)
	}

	return nil
}
