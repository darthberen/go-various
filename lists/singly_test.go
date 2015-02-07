package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushHead(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewSingly()
	for i, item := range data {
		list.PushHead(item)
		assert.NotNil(t, list.head, "head should contain a reference")
		assert.NotNil(t, list.tail, "head should contain a reference")
		assert.Exactly(t, item, list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(t, data[0], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func TestPushTail(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewSingly()
	for i, item := range data {
		list.PushTail(item)
		assert.NotNil(t, list.head, "head should contain a reference")
		assert.NotNil(t, list.tail, "head should contain a reference")
		assert.Exactly(t, data[0], list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(t, data[i], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func TestPopHead(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewSingly()
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

func TestPopTail(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewSingly()
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

func TestContains(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewSingly()
	_, err := list.Contains(data[0])
	assert.Exactly(t, ItemNotFoundError(data[0]), err, "wanted ItemNotFoundError")
	for _, item := range data {
		list.PushHead(item)
	}
	_, err = list.Contains("nonexistent data")
	assert.Exactly(t, ItemNotFoundError("nonexistent data"), err, "wanted ItemNotFoundError")
	for i := 0; i < len(data); i++ {
		exists, err := list.Contains(data[i])
		assert.NoError(t, err, "list should contain data[%d] '%s'", i, data[i])
		assert.True(t, exists, "list item data[%d] is incorrect", i)
	}
}

func TestDelete(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world", "another thing"}
	list := NewSingly()
	err := list.Delete(data[0])
	assert.Exactly(t, ItemNotFoundError(data[0]), err, "wanted ItemNotFoundError")
	for _, item := range data {
		list.PushHead(item)
	}
	err = list.Delete("nonexistent data")
	assert.Exactly(t, ItemNotFoundError("nonexistent data"), err, "wanted ItemNotFoundError")
	err = list.Delete(data[len(data)-1])
	assert.NoError(t, err, "list should delete head data[%d] '%s'", len(data)-1, data[len(data)-1])
	assert.Exactly(t, data[len(data)-2], list.head.Data, "list head data[%d] is incorrect", len(data)-2)

	err = list.Delete(data[0])
	assert.NoError(t, err, "list should delete tail data[%d] '%s'", 0, data[0])
	assert.Exactly(t, data[1], list.tail.Data, "list tail data[%d] is incorrect", len(data)-2)
	for i := 1; i < len(data)-1; i++ {
		err := list.Delete(data[i])
		assert.NoError(t, err, "list should delete data[%d] '%s'", i, data[i])
	}
}

func TestSize(t *testing.T) {
	data := []string{"data1", "data2", "something", "hello there world", "another thing"}
	list := NewSingly()
	size := list.Size()
	assert.Equal(t, 0, size, "expected list to be of size 0")
	for i, item := range data {
		list.PushHead(item)
		size := list.Size()
		assert.Equal(t, i+1, size, "PushHead: expected list to be of size %d", i)
	}
	for i := len(data) - 1; i >= 0; i-- {
		list.PopHead()
		size := list.Size()
		assert.Equal(t, i, size, "PopHead: expected list to be of size %d", i)
	}
	for i, item := range data {
		list.PushTail(item)
		size := list.Size()
		assert.Equal(t, i+1, size, "PushTail: expected list to be of size %d", i)
	}
	for i := len(data) - 1; i >= 0; i-- {
		list.PopTail()
		size := list.Size()
		assert.Equal(t, i, size, "PopTail: expected list to be of size %d", i)
	}
	for _, item := range data {
		list.PushHead(item)
	}
	for i := len(data) - 1; i >= 0; i-- {
		list.Delete(data[i])
		size := list.Size()
		assert.Equal(t, i, size, "Delete: expected list to be of size %d", i)
	}
}

func TestIsEmpty(t *testing.T) {
	data := "data1"
	list := NewSingly()
	assert.True(t, list.IsEmpty(), "initialized")
	list.PushHead(data)
	assert.False(t, list.IsEmpty(), "PushHead")
	list.PopHead()
	assert.True(t, list.IsEmpty(), "PopHead")
	list.PushTail(data)
	assert.False(t, list.IsEmpty(), "PushTail")
	list.PopTail()
	assert.True(t, list.IsEmpty(), "PopTail")
	list.PushTail(data)
	list.Delete(data)
	assert.True(t, list.IsEmpty(), "Delete")
}
