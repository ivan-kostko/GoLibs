//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package Map

import (
	"sync"
)

// The ThreadSafeMap type represents light weight and simple API for thread safe map
//
// NOTE(x): In case of operating on big amounts of data or need of extended functionality - consider to use https://github.com/streamrail/concurrent-map
type ThreadSafeMap struct {
	items map[string]interface{}
	sync.RWMutex
}

// Generic ThreadSafeMap factory
func New(initCap int) *ThreadSafeMap {
	items := make(map[string]interface{}, initCap)
	return &ThreadSafeMap{items: items}
}

// Wraps the map as ThreadSafe.
// > You should not use that map as general one anymore. Or just use MakeThreadSafeCopy
func MakeMapThreadSafe(m map[string]interface{}) *ThreadSafeMap {
	return &ThreadSafeMap{items: m}
}

// Makes ThreadSafe copy of the map.
func MakeThreadSafeCopy(m map[string]interface{}) *ThreadSafeMap {
	items := make(map[string]interface{}, len(m))
	for key, value := range m {
		items[key] = value
	}
	return &ThreadSafeMap{items: items}
}

// Makes ThreadSafe copy of the map recursively.
// In case the value is map[string]interface{} - it converts it into ThreadSafeMap recursively as well.
func MakeRecursivelyThreadSafeCopy(m map[string]interface{}) *ThreadSafeMap {
	items := make(map[string]interface{}, len(m))
	for key, value := range m {
		if x, ok := (value).(map[string]interface{}); ok {
			items[key] = MakeRecursivelyThreadSafeCopy(x)
		} else {
			items[key] = value
		}
	}
	return &ThreadSafeMap{items: items}
}

// Retrieves an element from map under given key.
func (tsm *ThreadSafeMap) Get(key string) (interface{}, bool) {
	tsm.RLock()
	defer tsm.RUnlock()

	val, ok := tsm.items[key]
	return val, ok
}

// Sets the given value under the specified key.
func (tsm *ThreadSafeMap) Set(key string, val interface{}) {
	tsm.Lock()
	defer tsm.Unlock()

	tsm.items[key] = val
}

// Sets the given value under the specified key and returns true, if the key didn't exist uppon invokation.
// Returns false and does nothing, in case there is already an entry with the same key
func (tsm *ThreadSafeMap) SetIfNotExists(key string, val interface{}) bool {
	tsm.Lock()
	defer tsm.Unlock()

	if _, ok := tsm.items[key]; !ok {
		tsm.items[key] = val
		return true
	}
	return false
}

// Removes an element from the map.
func (tsm *ThreadSafeMap) Remove(key string) {
	tsm.Lock()
	defer tsm.Unlock()

	delete(tsm.items, key)
}

// Returns copy of content as non thread safe map
func (tsm *ThreadSafeMap) Items() map[string]interface{} {
	tsm.RLock()
	defer tsm.RUnlock()
	x := make(map[string]interface{}, len(tsm.items))
	for key, value := range tsm.items {
		x[key] = value
	}
	return x
}
