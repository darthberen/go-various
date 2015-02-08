package lists

import "fmt"

// EmptyListError indicates that the request operation can't be performed as
// the list is empty
type EmptyListError string

func (e EmptyListError) Error() string {
	return fmt.Sprintf("lists: %s", e)
}
