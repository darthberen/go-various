package lists

import "sync"

// Singly goroutine-safe implementation of a singly-linked list
type Singly struct {
	head   *singlyNode
	tail   *singlyNode
	size   int64
	rwLock *sync.RWMutex
}

type singlyNode struct {
	Next *singlyNode
	Data string
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

// Size of the list O(1)
func (s *Singly) Size() int64 {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.size
}

// IsEmpty true if the singly-linked list contains no items O(1)
func (s *Singly) IsEmpty() bool {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.head == nil
}

// PushHead add data to the front of the list O(1)
func (s *Singly) PushHead(data string) {
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

// PushTail add data to the back of the list O(1)
func (s *Singly) PushTail(data string) {
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

// PopHead remove data from the front of the list O(1)
func (s *Singly) PopHead() (data string, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		return "", EmptyListError("can't remove an item from an empty list")
	}
	data = s.head.Data
	s.head = s.head.Next
	s.size--
	return
}

// PopTail remove data from the back of the list O(n)
func (s *Singly) PopTail() (data string, err error) {
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

// Contains data in the list, returns a ItemNotFoundError if the item
// does not exist O(n)
func (s *Singly) Contains(data string) (bool, error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	for tmp := s.head; tmp != nil; tmp = tmp.Next {
		if tmp.Data == data {
			return true, nil
		}
	}
	return false, ItemNotFoundError(data)
}

// Delete data in the list, returns a ItemNotFoundError if the item
// does not exist O(n)
func (s *Singly) Delete(data string) error {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if s.head == nil {
		return ItemNotFoundError(data)
	}
	if s.head.Data == data {
		if s.head == s.tail {
			s.head, s.tail = nil, nil
			s.size--
			return nil
		}
		s.head = s.head.Next
		s.size--
		return nil
	}
	for pred, tmp := s.head, s.head.Next; tmp != nil; pred, tmp = tmp, tmp.Next {
		if tmp.Data == data {
			pred.Next = tmp.Next
			if pred.Next == nil {
				s.tail = pred
			}
			s.size--
			return nil
		}
	}
	return ItemNotFoundError(data)
}

func (s *Singly) String() (str string) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	for tmp := s.head; tmp != nil; tmp = tmp.Next {
		str += tmp.Data
		if tmp.Next != nil {
			str += ", "
		}
	}
	return
}
