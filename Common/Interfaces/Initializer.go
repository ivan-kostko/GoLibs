package Interfaces

import (
	. "CustomErrors"
)

// Interface represents Initialize methood
type Initializer interface {
	// The method sets up instance and returns error if instance won't be initialized
	Initialize() *Error
}
