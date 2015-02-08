package lists

import "sync"

// Singly goroutine-safe implementation of a singly-linked list
type Singly struct {
	head   *singlyNode
	tail   *singlyNode
	size   int
	rwLock *sync.RWMutex
}

type singlyNode struct {
	Next *singlyNode
	Data interface{}
}

// NewSingly creates a new empty singly-linked list
func NewSingly() *Singly {
	return &Singly{
		head:   nil,
		tail:   nil,
		size:   0,
		rwLock: &sync.RWMutex{},
	}
}

// Size of the list
//
// Runtime: O(1)
func (s *Singly) Size() int {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.size
}

// IsEmpty returns true if the list contains no items
//
// Runtime: O(1)
func (s *Singly) IsEmpty() bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.head == nil
}

// PushHead adds data to the front of the list
//
// Runtime: O(1)
func (s *Singly) PushHead(data interface{}) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		s.head = &singlyNode{Next: nil, Data: data}
		s.tail = s.head
	} else {
		next := s.head
		s.head = &singlyNode{Next: next, Data: data}
	}
	s.size++
}

// PushTail adds data to the back of the list
//
// Runtime: O(1)
func (s *Singly) PushTail(data interface{}) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		s.head = &singlyNode{Next: nil, Data: data}
		s.tail = s.head
	} else {
		s.tail.Next = &singlyNode{Next: nil, Data: data}
		s.tail = s.tail.Next
	}
	s.size++
}

// PopHead removes data from the front of the list.  Returns an
// EmptyListError if there are no items in the list.
//
// Runtime: O(1)
func (s *Singly) PopHead() (data interface{}, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		return "", EmptyListError("can't remove an item from an empty list")
	}
	data = s.head.Data
	s.head = s.head.Next
	if s.head == nil {
		s.tail = nil
	}
	s.size--
	return
}

// PopTail removes data from the back of the list.  Returns an
// EmptyListError if there are no items in the list.
//
// Runtime: O(n)
func (s *Singly) PopTail() (data interface{}, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	var tmp *singlyNode
	if s.head == nil {
		return "", EmptyListError("can't remove an item from an empty list")
	}
	data = s.tail.Data
	if s.head == s.tail {
		s.head, s.tail = nil, nil
	} else {
		for tmp = s.head; tmp.Next != s.tail; tmp = tmp.Next {
		}
		s.tail = tmp
	}
	s.size--
	return
}

// Contains returns true if list contains any data where the comparison
// function returns true.  Moves from the head of the list to the tail.
//
// Runtime: O(n)
func (s *Singly) Contains(comparison func(data interface{}) (exists bool)) bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	for tmp := s.head; tmp != nil; tmp = tmp.Next {
		if comparison(tmp.Data) {
			return true
		}
	}
	return false
}

// Delete numItems data in the list based on the provided comparison function.  Moves from
// the head of list to the tail.  If the
// comparison function returns true for any item in the list then that item is
// deleted.  Returns the number of items that were deleted.  If numItems is <= 0 then all
// data in the list is scanned.
//
// Runtime: O(n)
func (s *Singly) Delete(numItems int, comparison func(data interface{}) (shouldDelete bool)) (numDeleted int) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		return
	}
	for comparison(s.head.Data) {
		numDeleted++
		s.size--
		if s.head == s.tail {
			s.head, s.tail = nil, nil
			return
		}
		s.head = s.head.Next
		if numItems == numDeleted {
			return
		}
	}
	for pred, tmp := s.head, s.head.Next; tmp != nil; pred, tmp = tmp, tmp.Next {
		if comparison(tmp.Data) {
			pred.Next = tmp.Next
			if pred.Next == nil {
				s.tail = pred
			}
			s.size--
			numDeleted++
			if numItems == numDeleted {
				return
			}
		}
	}
	return
}
