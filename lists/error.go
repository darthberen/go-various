package lists

import "fmt"

// EmptyListError indicates that the request operation can't be performed as
// the list is empty
type EmptyListError string

func (e EmptyListError) Error() string {
	return fmt.Sprintf("lists: %s", e)
}

// ItemNotFoundError indicates that the data is not contained within the list
type ItemNotFoundError string

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("lists: data '%s' could not be found", e)
}
