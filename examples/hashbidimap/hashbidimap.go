// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/rahul1534/gods-generic/flatten"

// HashBidiMapExample to demonstrate basic usage of HashMap
func main() {
	m := flatten.NewHashBidiMap[int, string]() // empty
	m.Put(1, "x")                              // 1->x
	m.Put(3, "b")                              // 1->x, 3->b (random order)
	m.Put(1, "a")                              // 1->a, 3->b (random order)
	m.Put(2, "b")                              // 1->a, 2->b (random order)
	_, _ = m.GetKey("a")                       // 1, true
	_, _ = m.Get(2)                            // b, true
	_, _ = m.Get(3)                            // nil, false
	_ = m.Values()                             // []interface {}{"a", "b"} (random order)
	_ = m.Keys()                               // []interface {}{1, 2} (random order)
	m.Remove(1)                                // 2->b
	m.Clear()                                  // empty
	m.Empty()                                  // true
	m.Size()                                   // 0
}
