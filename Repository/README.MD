# Repository
--
Repository project Repository.go

#### Install

    go get "github.com/ivan-kostko/GoLibs/Repository"

#### Import

    "github.com/ivan-kostko/GoLibs/Repository"

#### Description

    The package provides generic interface to communicate beween business and persistence domains.

### Usage

```go
const (
	ERR_FAILEDTOGENERATEINSTRUCTION = "Repository: Failed to generate instruction"
	ERR_FAILEDTOEXECUTEINSTRUCTION  = "Repository: Failed to execute instruction"
	ERR_FAILEDTOINTERPRETDSRESULT   = "Repository: Failed to interpret data source result"
)
```

##### type DomainModel

```go
type DomainModel interface{}
```

An alias for models in domain

##### type FilteringCondition

```go
type FilteringCondition map[string][]Predicate
```

Represents filtering creteria, which is applied at DataSource instruction.
For.Ex. for SQL DataSource it would be converted by Instructor into WHERE
clause. It is implemented as composition of fields(key(s)) and predicates(value)

##### type Instructor

```go
type Instructor interface {
	// Translates FilteringCondition(s) into data source instruction
	GenerateInstruction(...FilteringCondition) (ds.Instruction, *Error)
}
```

Represents translator responsible for generating DataSource.Instruction

##### type Operator

```go
type Operator string
```

Represents comparison operator ENUM

```go
const (
	NOTIN Operator = "NOTIN" // NOT in range
	IN    Operator = "IN"    // In range
	GE    Operator = "GE"    // Greater or equal
	GT    Operator = "GT"    // Greater than
	LE    Operator = "LE"    // Less or equal
	LT    Operator = "LT"    // Less than
)
```

##### type Predicate

```go
type Predicate struct {
	Operator
	Values []interface{}
}
```

Represents single predicate as operator + value(s)

##### type Repository

```go
type Repository struct {
}
```

Represents a data repository abstraction. It is supposed to be an interface to
request and retreive entities between domains.

##### func  GetNewRepository

```go
func GetNewRepository(dataSource *ds.DataSource, instructor Instructor, resultInterpreter ResultInterpreter, checkError func(*Error) bool) *Repository
```
Gets complete repository

##### func (*Repository) GetAll

```go
func (rep *Repository) GetAll(fc ...FilteringCondition) ([]DomainModel, *Error)
```
Retreives all entities conforming FilteringCondition(s)

##### type ResultInterpreter

```go
type ResultInterpreter interface {
	Interpret(ds.Result) ([]DomainModel, *Error)
}
```

Represents service translating datasource result into slice of complete domain
entities
