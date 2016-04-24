# Interfaces
--
    import "."

### Import

    "github.com/ivan-kostko/GoLibs/Common/Interfaces"

### Install

    go get github.com/ivan-kostko/GoLibs/Common/Interfaces

## Usage

#### type Disposer

    type Disposer interface {
    	// The method "cleans" all internal references to let current instance to be garbage collected
    	Dispose()
    }


Interface represents Dispose method

#### type Initializer

    type Initializer interface {
    	// The method sets up instance and returns error if instance won't be initialized
    	Initialize() *Error
    }


Interface represents Initialize methood

#### type MustInitializer

    type MustInitializer interface {
    	// The method sets up instance and panics if instance won't be initialized
    	MustInitialize()
    }


Interface represents MustInitialize methood
