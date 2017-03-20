# Variables
--
    import "GoLibs/Variables"

### Description

The package provides unified interface to access variables. Under the hood it
could use Consul, Vault, etcd, etc. - for developer it should not change
anything.

## Usage

#### type RWLock

```go
type RWLock interface {
	sync.Locker
	RLock()
	RUnlock()
}
```


#### type Trigger

```go
type Trigger func(v *Variable)
```


#### type Variable

```go
type Variable struct {
}
```

Represents generic variable.

By design it supposed to be a thread safe.

#### func  New

```go
func New() *Variable
```
Variable factory.

#### func (*Variable) HookupAfterLocalChangeTrigger

```go
func (this *Variable) HookupAfterLocalChangeTrigger(callback Trigger)
```
Hookup `callback` with after local change event, fired by
[SetValue](#func-variable-setvalue).

#### func (*Variable) InterfaceValue

```go
func (this *Variable) InterfaceValue() interface{}
```
Returns variable current value as `interface{}`.

#### func (*Variable) SetValue

```go
func (this *Variable) SetValue(i interface{})
```
Variable value setter. In case `i` is not equal to original variable value - it
assigns `i` value to variable and invokes afterLocalChange hooked triggers.
