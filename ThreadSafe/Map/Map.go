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

// The ThreadSafeMap type represents represents light weight and simple API for thread safe map
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

// Retrieves an element from map under given key.
func (tsm ThreadSafeMap) Get(key string) (interface{}, bool) {
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

// Removes an element from the map.
func (tsm *ThreadSafeMap) Remove(key string) {
	tsm.Lock()
	defer tsm.Unlock()

	delete(tsm.items, key)
}
