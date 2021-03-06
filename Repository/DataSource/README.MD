# DataSource
--
DataSource project DataSource.go

#### Install

    go get "github.com/ivan-kostko/GoLibs/Repository/DataSource"

#### Import

    "github.com/ivan-kostko/GoLibs/Repository/DataSource"

#### Description

    The package provides generic interface to access persistence.

### Usage

##### type DataSource

```go
type DataSource struct {
	ExecuteInstruction
}
```

Represents generic data source or persistence

##### func  GetNewDataSource

```go
func GetNewDataSource(ei ExecuteInstruction) *DataSource
```
Represents generic DataSource factory

##### type ExecuteInstruction

```go
type ExecuteInstruction func(i Instruction) (Result, *Error)
```

Represents standard instruction execution function

##### type Instruction

```go
type Instruction string
```

Represents the instruction to DataSource For SQL DataSource it would be a
SQL-script; for http service DataSource it would be an access link; and so on

##### type Result

```go
type Result interface{}
```

Represents a result of the instruction
