package gohilbert

import "errors"

// Errors returned when validating input.
var (
	ErrLevelOutOfRange = errors.New("Level is out of range")
	ErrXCoordinateOutOfRange = errors.New("X coordinate is out of range")
	ErrYCoordinateOutOfRange = errors.New("Y coordinate is out of range")
	ErrIndexOutOfRange = errors.New("Index is out of range")
)
