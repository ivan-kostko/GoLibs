# Map
##### Import

    "github.com/ivan-kostko/ThreadSafe/Map"

##### Install

    go get github.com/ivan-kostko/ThreadSafe/Map


##### Description

The ThreadSafeMap package represents light weight and simple API for thread safe
map. The functionality is very limited to make it clean and robust. In case of
operating on big amounts of data or need of extended functionality - consider to
use https://github.com/streamrail/concurrent-map

#### Usage

##### type ThreadSafeMap

```go
type ThreadSafeMap struct {
	sync.RWMutex
}
```

The ThreadSafeMap type represents light weight and simple API for thread safe
map

NOTE(x): In case of operating on big amounts of data or need of extended
functionality - consider to use https://github.com/streamrail/concurrent-map

--
##### func  MakeMapThreadSafe

```go
func MakeMapThreadSafe(m map[string]interface{}) *ThreadSafeMap
```
Wraps the map as ThreadSafe.

> You should not use that map as general one anymore. Or just use
MakeThreadSafeCopy

--
##### func  MakeRecursivelyThreadSafeCopy

```go
func MakeRecursivelyThreadSafeCopy(m map[string]interface{}) *ThreadSafeMap
```
Makes ThreadSafe copy of the map recursively. In case the value is
map[string]interface{} - it converts it into ThreadSafeMap recursively as well.

--
##### func  MakeThreadSafeCopy

```go
func MakeThreadSafeCopy(m map[string]interface{}) *ThreadSafeMap
```
Makes ThreadSafe copy of the map.

--
##### func  New

```go
func New(initCap int) *ThreadSafeMap
```
Generic ThreadSafeMap factory

--
##### func (*ThreadSafeMap) Get

```go
func (tsm *ThreadSafeMap) Get(key string) (interface{}, bool)
```
Retrieves an element from map under given key. Returns false in case there is no
entry associated with the key

--
##### func (*ThreadSafeMap) Items

```go
func (tsm *ThreadSafeMap) Items() map[string]interface{}
```
Returns copy of content as non thread safe map

--
##### func (*ThreadSafeMap) Remove

```go
func (tsm *ThreadSafeMap) Remove(key string)
```
Removes an element from the map.

--
##### func (*ThreadSafeMap) Set

```go
func (tsm *ThreadSafeMap) Set(key string, val interface{})
```
Sets the given value under the specified key.

--
##### func (*ThreadSafeMap) SetIfNotExists

```go
func (tsm *ThreadSafeMap) SetIfNotExists(key string, val interface{}) bool
```
Sets the given value under the specified key and returns true, if the key didn't
exist uppon invokation. Returns false and does nothing, in case there is already
an entry with the same key

--
