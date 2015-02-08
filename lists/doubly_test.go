package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyPushHead(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	for i, item := range data {
		list.PushHead(item)
		assert.NotNil(t, list.head, "head should contain a reference")
		assert.NotNil(t, list.tail, "head should contain a reference")
		assert.Exactly(t, item, list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(t, data[0], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func TestDoublyPushTail(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	for i, item := range data {
		list.PushTail(item)
		assert.NotNil(t, list.head, "head should contain a reference")
		assert.NotNil(t, list.tail, "head should contain a reference")
		assert.Exactly(t, data[0], list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(t, data[i], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func TestDoublyPopHead(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	_, err := list.PopHead()
	assert.Exactly(t, EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range data {
		list.PushHead(item)
	}
	for i := len(data) - 1; i >= 0; i-- {
		item, err := list.PopHead()
		assert.NoError(t, err, "list should contain items")
		assert.Exactly(t, data[i], item, "list item data[%d] is incorrect", i)
	}
}

func TestDoublyPopTail(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	_, err := list.PopTail()
	assert.Exactly(t, EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range data {
		list.PushHead(item)
	}
	for i := 0; i < len(data); i++ {
		item, err := list.PopTail()
		assert.NoError(t, err, "list should contain items")
		assert.Exactly(t, data[i], item, "list item data[%d] is incorrect", i)
	}
}
