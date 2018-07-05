// Copyright 2017 Annchain Information Technology Services Co.,Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package mlist

type Element struct {
	key   string
	value interface{}
	next  *Element
	prev  *Element
}

func (e *Element) clear() {
	e.next = nil
	e.prev = nil
}

type MapList struct {
	head  *Element
	tail  *Element
	keyMp map[string]*Element
}

func NewMapList() *MapList {
	var m MapList
	m.Init()
	return &m
}

func (l *MapList) Init() {
	l.keyMp = make(map[string]*Element)
}

func (l *MapList) Len() int {
	return len(l.keyMp)
}

// Set if the key can't find in the list
// then insert the value into the end of the list,
// or modify the value of the key
func (l *MapList) Set(key string, value interface{}) {
	if v, ok := l.keyMp[key]; ok {
		v.value = value
		return
	}
	e := &Element{
		key: key,
	}
	e.value = value
	if l.tail != nil {
		l.tail.next = e
	}
	e.prev = l.tail
	l.tail = e
	if l.head == nil {
		l.head = e
	}
	l.keyMp[key] = e
}

// Get get the key in the list
func (l *MapList) Get(key string) (interface{}, bool) {
	if v, ok := l.keyMp[key]; ok {
		return v.value, ok
	}
	return nil, false
}

// Del delete the key in the list
func (l *MapList) Del(key string) {
	if v, ok := l.keyMp[key]; ok {
		if l.head == v {
			l.head = v.next
		}
		if l.tail == v {
			l.tail = v.prev
		}
		if v.prev != nil {
			v.prev.next = v.next
		}
		if v.next != nil {
			v.next.prev = v.prev
		}
		v.clear()
		delete(l.keyMp, key)
	}
}

// Has check whether the key is in the list
func (l *MapList) Has(key string) bool {
	_, ok := l.keyMp[key]
	return ok
}

// Exec go through the list
func (l *MapList) Exec(exec func(string, interface{})) {
	var p *Element
	for p = l.head; p != nil; p = p.next {
		exec(p.key, p.value)
	}
}

// Exec go through the list
func (l *MapList) ExecBreak(exec func(string, interface{}) bool) bool {
	var p *Element
	for p = l.head; p != nil; p = p.next {
		if !exec(p.key, p.value) {
			return false
		}
	}
	return true
}

func (l *MapList) Reset() {
	l.Init()
	l.head = nil
	l.tail = nil
}
