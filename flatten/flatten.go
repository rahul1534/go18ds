package flatten

import (
	"github.com/rahul1534/gods-generic/lists/arraylist"
	"github.com/rahul1534/gods-generic/lists/doublylinkedlist"
	"github.com/rahul1534/gods-generic/lists/singlylinkedlist"
)

// #region List implementations
func NewArrayList[T comparable]() *arraylist.List[T] {
	return arraylist.New[T]()
}

func NewSinglyLinkedList[T comparable]() *singlylinkedlist.List[T] {
	return singlylinkedlist.New[T]()
}

func NewDoublyLinkedList[T comparable]() *doublylinkedlist.List[T] {
	return doublylinkedlist.New[T]()
}

// #endregion
