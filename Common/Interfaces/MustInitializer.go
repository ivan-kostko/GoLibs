package Interfaces

// Interface represents MustInitialize methood
type MustInitializer interface {
	// The method sets up instance and panics if instance won't be initialized
	MustInitialize()
}
