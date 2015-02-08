package lists

import "sync"

// Doubly goroutine-safe implementation of a doubly-linked list
type Doubly struct {
	head   *doublyNode
	tail   *doublyNode
	size   int64
	rwLock *sync.RWMutex
}

type doublyNode struct {
	Next *doublyNode
	Prev *doublyNode
	Data string
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

// Size of the list O(1)
func (d *Doubly) Size() int64 {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	return d.size
}

// IsEmpty true if the list contains no items O(1)
func (d *Doubly) IsEmpty() bool {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	return d.head == nil
}

// PushHead add data to the front of the list O(1)
func (d *Doubly) PushHead(data string) {
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

// PushTail add data to the back of the list O(1)
func (d *Doubly) PushTail(data string) {
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

// PopHead remove data from the front of the list O(1)
func (d *Doubly) PopHead() (data string, err error) {
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

// PopTail remove data from the back of the list O(1)
func (d *Doubly) PopTail() (data string, err error) {
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

func (d *Doubly) String() (str string) {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	for tmp := d.head; tmp != nil; tmp = tmp.Next {
		str += tmp.Data
		if tmp.Next != nil {
			str += ", "
		}
	}
	return
}
