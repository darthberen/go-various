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

// Contains returns true if list contains any data where the comparison
// function returns true.  Moves from the head of the list to the tail.
//
// Runtime: O(n)
func (d *Doubly) Contains(comparison func(data interface{}) (exists bool)) bool {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	for tmp := d.head; tmp != nil; tmp = tmp.Next {
		if comparison(tmp.Data) {
			return true
		}
	}
	return false
}

// Delete numItems data in the list based on the provided comparison function.
// Moves from the head of list to the tail.  If the
// comparison function returns true for any item in the list then that item is
// deleted.  Returns the number of items that were deleted.  If numItems is <= 0
// then all data in the list is scanned.
//
// Runtime: O(n)
func (d *Doubly) Delete(numItems int, comparison func(data interface{}) (shouldDelete bool)) (numDeleted int) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	if d.head == nil {
		return
	}
	for tmp := d.head; tmp != nil; tmp = tmp.Next {
		if comparison(tmp.Data) {
			pred := tmp.Prev
			succ := tmp.Next
			if succ == nil {
				d.tail = pred
			} else {
				succ.Prev = pred
			}
			if pred == nil {
				d.head = succ
			} else {
				pred.Next = succ
			}
			d.size--
			numDeleted++
			if numItems == numDeleted {
				return
			}
		}
	}
	return
}
