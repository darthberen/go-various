package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DoublyTestSuite struct {
	suite.Suite
	findExact func(string) func(interface{}) bool
	data      []string
}

func TestDoublyTestSuite(t *testing.T) {
	suite.Run(t, new(DoublyTestSuite))
}

func (suite *DoublyTestSuite) SetupTest() {
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

func (suite *DoublyTestSuite) TestNewDoubly() {
	list := NewDoubly()
	assert.Nil(suite.T(), list.head)
	assert.Nil(suite.T(), list.tail)
	assert.Equal(suite.T(), 0, list.size)
}

func (suite *DoublyTestSuite) TestPushHead() {
	list := NewDoubly()
	for i, item := range suite.data {
		list.PushHead(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), item, list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), suite.data[0], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *DoublyTestSuite) TestPushTail() {
	list := NewDoubly()
	for i, item := range suite.data {
		list.PushTail(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), suite.data[0], list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), suite.data[i], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *DoublyTestSuite) TestPopHead() {
	list := NewDoubly()
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

func (suite *DoublyTestSuite) TestPopTail() {
	list := NewDoubly()
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

func (suite *DoublyTestSuite) TestContains() {
	list := NewDoubly()
	contains := list.Contains(suite.findExact(suite.data[0]))
	assert.False(suite.T(), contains, "empty list should not contain anything")
	for _, item := range suite.data {
		list.PushHead(item)
	}
	contains = list.Contains(suite.findExact("nonexistent data"))
	assert.False(suite.T(), contains)
	for i := 0; i < len(suite.data); i++ {
		contains = list.Contains(suite.findExact(suite.data[i]))
		assert.True(suite.T(), contains, "list item data[%d] is missing", i)
	}
}
