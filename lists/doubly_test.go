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
}

func (suite *DoublyTestSuite) TestDoublyPushHead() {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	for i, item := range data {
		list.PushHead(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), item, list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), data[0], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *DoublyTestSuite) TestDoublyPushTail() {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	for i, item := range data {
		list.PushTail(item)
		assert.NotNil(suite.T(), list.head, "head should contain a reference")
		assert.NotNil(suite.T(), list.tail, "head should contain a reference")
		assert.Exactly(suite.T(), data[0], list.head.Data, "head data[%d] is incorrect", i)
		assert.Exactly(suite.T(), data[i], list.tail.Data, "tail data[%d] is incorrect", i)
	}
}

func (suite *DoublyTestSuite) TestDoublyPopHead() {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	_, err := list.PopHead()
	assert.Exactly(suite.T(), EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range data {
		list.PushHead(item)
	}
	for i := len(data) - 1; i >= 0; i-- {
		item, err := list.PopHead()
		assert.NoError(suite.T(), err, "list should contain items")
		assert.Exactly(suite.T(), data[i], item, "list item data[%d] is incorrect", i)
	}
}

func (suite *DoublyTestSuite) TestDoublyPopTail() {
	data := []string{"data1", "data2", "something", "hello there world"}
	list := NewDoubly()
	_, err := list.PopTail()
	assert.Exactly(suite.T(), EmptyListError("can't remove an item from an empty list"), err, "wanted EmptyListError")
	for _, item := range data {
		list.PushHead(item)
	}
	for i := 0; i < len(data); i++ {
		item, err := list.PopTail()
		assert.NoError(suite.T(), err, "list should contain items")
		assert.Exactly(suite.T(), data[i], item, "list item data[%d] is incorrect", i)
	}
}
