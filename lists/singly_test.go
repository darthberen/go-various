package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SinglyTestSuite struct {
	suite.Suite
	findExact func(string) func(interface{}) bool
	data      []string
}

func TestSinglyTestSuite(t *testing.T) {
	suite.Run(t, new(SinglyTestSuite))
}

func (suite *SinglyTestSuite) SetupTest() {
	suite.data = []string{"data1", "data2", "something", "hello there world", "another thing"}
	suite.findExact = func(item string) func(interface{}) bool {
		return func(data interface{}) bool {
			if data.(string) == item {
				return true
			}
			return false
		}
	}
}

func (suite *SinglyTestSuite) TestNewSingly() {
	list := NewSingly()
	assert.Nil(suite.T(), list.head)
	assert.Nil(suite.T(), list.tail)
	assert.Equal(suite.T(), 0, list.size)
}

func (suite *SinglyTestSuite) TestSinglyPushHead() {
	list := NewSingly()
	for i, item := range suite.data {
		list.PushHead(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), item, list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), suite.data[0], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyPushTail() {
	list := NewSingly()
	for i, item := range suite.data {
		list.PushTail(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), suite.data[0], list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), suite.data[i], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyPopHead() {
	list := NewSingly()
	_, err := list.PopHead()
	assert.Exactly(suite.T(), EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range suite.data {
		list.PushHead(item)
	}
	for i := len(suite.data) - 1; i >= 0; i-- {
		item, err := list.PopHead()
		assert.NoError(suite.T(), err, "list should contain items")
		assert.Exactly(suite.T(), suite.data[i], item, "list item data[%d] is incorrect", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyPopTail() {
	list := NewSingly()
	_, err := list.PopTail()
	assert.Exactly(suite.T(), EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range suite.data {
		list.PushHead(item)
	}
	for i := 0; i < len(suite.data); i++ {
		item, err := list.PopTail()
		assert.NoError(suite.T(), err, "list should contain items")
		assert.Exactly(suite.T(), suite.data[i], item, "list item data[%d] is incorrect", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyContains() {
	list := NewSingly()
	contains := list.Contains(suite.findExact(suite.data[0]))
	assert.False(suite.T(), contains, "empty list should not contain anything")
	for _, item := range suite.data {
		list.PushHead(item)
	}
	contains = list.Contains(suite.findExact("nonexistent data"))
	assert.False(suite.T(), contains)
	for i := 0; i < len(suite.data); i++ {
		contains = list.Contains(suite.findExact(suite.data[i]))
		assert.True(suite.T(), contains, "list item data[%d] is incorrect", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyDelete() {
	list := NewSingly()
	numDeleted := list.Delete(0, suite.findExact(""))
	assert.Equal(suite.T(), 0, numDeleted, "expected 0 items to be deleted from an empty list")
	for _, item := range suite.data {
		list.PushHead(item)
	}
	numDeleted = list.Delete(0, suite.findExact("nonexistent data"))
	assert.Equal(suite.T(), 0, numDeleted, "expected 0 items to be deleted")
	numDeleted = list.Delete(0, suite.findExact(suite.data[len(suite.data)-1]))
	assert.Equal(suite.T(), 1, numDeleted, "list should delete head data[%d] '%s'", len(suite.data)-1, suite.data[len(suite.data)-1])
	assert.Exactly(suite.T(), suite.data[len(suite.data)-2], list.head.Data, "list head data[%d] is incorrect", len(suite.data)-2)

	numDeleted = list.Delete(0, suite.findExact(suite.data[0]))
	assert.Equal(suite.T(), 1, numDeleted, "list should delete tail data[%d] '%s'", 0, suite.data[0])
	assert.Exactly(suite.T(), suite.data[1], list.tail.Data, "list tail data[%d] is incorrect", len(suite.data)-2)
	for i := 1; i < len(suite.data)-1; i++ {
		numDeleted = list.Delete(0, suite.findExact(suite.data[i]))
		assert.Equal(suite.T(), 1, numDeleted, "list should delete data[%d] '%s'", i, suite.data[i])
	}

	for i := 0; i < 10; i++ {
		list.PushHead(i)
	}
	numDeleted = list.Delete(1, func(data interface{}) bool { return true })
	assert.Equal(suite.T(), 1, numDeleted, "list should delete 1 item")
	assert.Equal(suite.T(), 9, list.Size(), "list should have 9 items")
	numDeleted = list.Delete(10, func(data interface{}) bool { return data.(int) > 0 })
	assert.Equal(suite.T(), 8, numDeleted, "list should delete 8 items")
	assert.Equal(suite.T(), 1, list.Size(), "list should have 1 item")
	assert.Equal(suite.T(), 0, list.head.Data)
	assert.Exactly(suite.T(), list.head, list.tail)
}

func (suite *SinglyTestSuite) TestSinglySize() {
	list := NewSingly()
	size := list.Size()
	assert.Equal(suite.T(), 0, size, "expected list to be of size 0")
	for i, item := range suite.data {
		list.PushHead(item)
		size := list.Size()
		assert.Equal(suite.T(), i+1, size, "PushHead: expected list to be of size %d", i)
	}
	for i := len(suite.data) - 1; i >= 0; i-- {
		list.PopHead()
		size := list.Size()
		assert.Equal(suite.T(), i, size, "PopHead: expected list to be of size %d", i)
	}
	for i, item := range suite.data {
		list.PushTail(item)
		size := list.Size()
		assert.Equal(suite.T(), i+1, size, "PushTail: expected list to be of size %d", i)
	}
	for i := len(suite.data) - 1; i >= 0; i-- {
		list.PopTail()
		size := list.Size()
		assert.Equal(suite.T(), i, size, "PopTail: expected list to be of size %d", i)
	}
	for _, item := range suite.data {
		list.PushHead(item)
	}
	for i := len(suite.data) - 1; i >= 0; i-- {
		list.Delete(0, suite.findExact(suite.data[i]))
		size := list.Size()
		assert.Equal(suite.T(), i, size, "Delete: expected list to be of size %d", i)
	}
}

func (suite *SinglyTestSuite) TestSinglyIsEmpty() {
	data := "data1"
	list := NewSingly()
	assert.True(suite.T(), list.IsEmpty(), "initialized")
	list.PushHead(data)
	assert.False(suite.T(), list.IsEmpty(), "PushHead")
	list.PopHead()
	assert.True(suite.T(), list.IsEmpty(), "PopHead")
	list.PushTail(data)
	assert.False(suite.T(), list.IsEmpty(), "PushTail")
	list.PopTail()
	assert.True(suite.T(), list.IsEmpty(), "PopTail")
	list.PushTail(data)
	list.Delete(0, suite.findExact(data))
	assert.True(suite.T(), list.IsEmpty(), "Delete")
}
