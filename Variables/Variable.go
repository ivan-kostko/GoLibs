package Variables

import (
	"sync"
)

type Trigger func(v *Variable)

type RWLock interface {
	sync.Locker
	RLock()
	RUnlock()
}

// Represents generic variable.
//
// By design it supposed to be a thread safe.
type Variable struct {
	value             interface{}
	afterLocalChange  Trigger
	afterRemoteChange Trigger
	lock              RWLock
}

// Variable factory.
func New() *Variable {

	return &Variable{
		lock: &sync.RWMutex{},
	}
}

// Variable value setter. In case `i` is not equal to original variable value - it assigns `i` value to variable and invokes afterLocalChange hooked triggers.
func (this *Variable) SetValue(i interface{}) {
	if this.value != i {
		this.setValue(i)
		if this.afterLocalChange != nil {
			this.afterLocalChange(this)
		}
	}
}

// Returns variable current value as `interface{}`.
func (this *Variable) InterfaceValue() interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.value
}

// Hookup `callback` with after local change event, fired by [SetValue](#func-variable-setvalue).
func (this *Variable) HookupAfterLocalChangeTrigger(callback Trigger) {
	if callback != nil {
		fn := *new(Trigger)

		if this.afterLocalChange != nil {
			oldTrigger := this.afterLocalChange
			fn = func(v *Variable) {
				oldTrigger(v)
				callback(v)
			}
		} else {
			fn = func(v *Variable) {
				callback(v)
			}
		}
		this.afterLocalChange = fn

	}
}

// Aquires a lock and assigns `i` to Variables value.
// NB: Does not fire any trigger, and still will assign `i` even it is equal(`==`) to original value.
func (this *Variable) setValue(i interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = i
}
