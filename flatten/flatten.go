package flatten

import (
	"github.com/rahul1534/gods-generic/lists/arraylist"
	"github.com/rahul1534/gods-generic/lists/doublylinkedlist"
	"github.com/rahul1534/gods-generic/lists/singlylinkedlist"
	"github.com/rahul1534/gods-generic/maps/hashbidimap"
	"github.com/rahul1534/gods-generic/maps/hashmap"
	"github.com/rahul1534/gods-generic/maps/linkedhashmap"
	"github.com/rahul1534/gods-generic/maps/treebidimap"
	"github.com/rahul1534/gods-generic/maps/treemap"
	"github.com/rahul1534/gods-generic/sets/hashset"
	"github.com/rahul1534/gods-generic/sets/linkedhashset"
	"github.com/rahul1534/gods-generic/sets/treeset"
	"github.com/rahul1534/gods-generic/stacks/arraystack"
	"github.com/rahul1534/gods-generic/stacks/linkedliststack"
	"github.com/rahul1534/gods-generic/trees/avltree"
	"github.com/rahul1534/gods-generic/trees/binaryheap"
	"github.com/rahul1534/gods-generic/trees/btree"
	"github.com/rahul1534/gods-generic/trees/redblacktree"
	"github.com/rahul1534/gods-generic/utils"
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

// #region Map implementations
func NewHashMap[K comparable, V comparable]() *hashmap.Map[K, V] {
	return hashmap.New[K, V]()
}

func NewLinkedHashMap[K comparable, V comparable]() *linkedhashmap.Map[K, V] {
	return linkedhashmap.New[K, V]()
}

func NewTreeMap[K comparable, V comparable](comparator utils.Comparator[K]) *treemap.Map[K, V] {
	return treemap.NewWith[K, V](comparator)
}

func NewHashBidiMap[K comparable, V comparable]() *hashbidimap.Map[K, V] {
	return hashbidimap.New[K, V]()
}

func NewTreeBidiMap[K comparable, V comparable](keyComparator utils.Comparator[K], valueComparator utils.Comparator[V]) *treebidimap.Map[K, V] {
	return treebidimap.NewWith(keyComparator, valueComparator)
}

// #endregion

// #region Set implementations
func NewHashSet[T comparable]() *hashset.Set[T] {
	return hashset.New[T]()
}

func NewLinkedHashSet[T comparable]() *linkedhashset.Set[T] {
	return linkedhashset.New[T]()
}

func NewTreeSet[T comparable](comparator utils.Comparator[T]) *treeset.Set[T] {
	return treeset.NewWith(comparator)
}

// #endregion

// #region Stack implementations
func NewArrayStack[T comparable]() *arraystack.Stack[T] {
	return arraystack.New[T]()
}

func NewLinkedListStack[T comparable]() *linkedliststack.Stack[T] {
	return linkedliststack.New[T]()
}

// #endregion

// #region Tree implementations
func NewAvlTree[K comparable, V comparable](comparator utils.Comparator[K]) *avltree.Tree[K, V] {
	return avltree.NewWith[K, V](comparator)
}

func NewBinaryHeap[T comparable](comparator utils.Comparator[T]) *binaryheap.Heap[T] {
	return binaryheap.NewWith(comparator)
}

func NewBtree[K comparable, V comparable](order int, comparator utils.Comparator[K]) *btree.Tree[K, V] {
	return btree.NewWith[K, V](order, comparator)
}

func NewRedBlackTree[K comparable, V comparable](comparator utils.Comparator[K]) *redblacktree.Tree[K, V] {
	return redblacktree.NewWith[K, V](comparator)
}

// #endregion
