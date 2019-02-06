package gengou

import "time"

// TA is a variable for TimeAdapter to be injected
var TA TimeAdapter = &LocalTimeAdapter{}

// TimeAdapter is an interface for inject time package
type TimeAdapter interface {
	LoadLocation(name string) (*time.Location, error)
}

// LocalTimeAdapter is implementation for TimeAdapter
type LocalTimeAdapter struct {
}

// LoadLocation is a func for TimeAdapter, actually, this func is just an adapter to time package
func (l *LocalTimeAdapter) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
