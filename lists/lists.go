/*
Package lists provides an implementation of goroutine (thread) safe lists.
It uses a RWMutex (http://golang.org/pkg/sync/#RWMutex) to manage access
to the lists.

Note that you probably don't need to use this package if you are looking for
a queue (FIFO, first-in first-out) data structure.  You can use a channel instead.
This package could be useful for LIFO, last-in last-out, (singly-linked list)
or LIFO and FIFO functionality (doubly-linked list) when it is required.  I would
recommend reading http://blog.golang.org/share-memory-by-communicating before using
this package if you have not.

Submit any issues or feature requests here: https://github.com/suicidejack/go-various/issues
*/
package lists
