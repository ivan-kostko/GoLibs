package Interfaces

// Interface represents Dispose method
type Disposer interface {
	// The method "cleans" all internal references to let current instance to be garbage collected
	Dispose()
}
