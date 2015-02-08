package lists

import "sync"

// Doubly goroutine-safe implementation of a doubly-linked list
type Doubly struct {
	head   *doublyNode
	tail   *doublyNode
	size   int
	rwLock *sync.RWMutex
}

type doublyNode struct {
	Next *doublyNode
	Prev *doublyNode
	Data interface{}
}

// NewDoubly creates a new empty doubly-linked list
func NewDoubly() *Doubly {
	return &Doubly{
		head:   nil,
		tail:   nil,
		size:   0,
		rwLock: &sync.RWMutex{},
	}
}

// Size of the list
//
// Runtime: O(1)
func (d *Doubly) Size() int {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	return d.size
}

// IsEmpty true if the list contains no items
//
// Runtime: O(1)
func (d *Doubly) IsEmpty() bool {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	return d.head == nil
}

// PushHead adds data to the front of the list
//
// Runtime: O(1)
func (d *Doubly) PushHead(data interface{}) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	if d.head == nil {
		d.head = &doublyNode{Next: nil, Prev: nil, Data: data}
		d.tail = d.head
	} else {
		d.head.Prev = &doublyNode{Next: d.head, Prev: nil, Data: data}
		d.head = d.head.Prev
	}
	d.size++
}

// PushTail adds data to the back of the list
//
// Runtime: O(1)
func (d *Doubly) PushTail(data interface{}) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	if d.head == nil {
		d.head = &doublyNode{Next: nil, Prev: nil, Data: data}
		d.tail = d.head
	} else {
		d.tail.Next = &doublyNode{Next: nil, Prev: d.tail, Data: data}
		d.tail = d.tail.Next
	}
	d.size++
}

// PopHead removes data from the front of the list.  Returns an
// EmptyListError if there are no items in the list.
//
// Runtime: O(1)
func (d *Doubly) PopHead() (data interface{}, err error) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	if d.head == nil {
		return "", EmptyListError("can't remove an item from an empty list")
	}
	data = d.head.Data
	d.head = d.head.Next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.Prev = nil
	}
	d.size--
	return
}

// PopTail removes data from the back of the list.  Returns an
// EmptyListError if there are no items in the list.
//
// Runtime: O(1)
func (d *Doubly) PopTail() (data interface{}, err error) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	if d.head == nil {
		return "", EmptyListError("can't remove an item from an empty list")
	}
	data = d.tail.Data
	if d.head == d.tail {
		d.head, d.tail = nil, nil
	} else {
		d.tail = d.tail.Prev
		d.tail.Next = nil
	}
	d.size--
	return
}
