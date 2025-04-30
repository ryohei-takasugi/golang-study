package error

import (
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) String() string {
	return "at " + e.When.Format(time.RFC3339) + ", " + e.What
}
