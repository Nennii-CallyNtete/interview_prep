// Data Structures in Golang
// They are as follws:
// Runes
// Strings
// Arrays
// Slices
// Maps
// Structs
// Interfaces
// LL
// DLL
// Graphs
// Trees

package main

import "fmt"

// Structs. Golang equivalent of JS Objects.
type Test struct {
	Score   int64
	Comment string
	Owner   string
}

// Linked Lists by Nennster.

type LL struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Value int64
	Next  *Node
}

func (L *LL) push(node *Node) error {
	if L.Length == 0 {
		L.Head = node
		L.Tail = node
	} else {
		tail := L.Tail
		tail.Next = node
		L.Tail = tail.Next
		L.Tail.Next = nil
	}

	L.Length += 1
	return nil
}

func (L *LL) pop() *Node {
	if L.Length == 0 {
		return nil
	}
	var temp = L.Head

	for i := 1; i < L.Length-1; i++ {
		temp = temp.Next
	}
	value := temp.Next
	temp.Next = nil
	L.Tail = temp
	L.Length -= 1

	return value
}

func (L *LL) get(n int) *Node {
	if L.Length == 0 {
		return nil
	}
	if n >= L.Length || n < 0 {
		return nil
	}

	var temp = L.Head

	for i := 0; i < n; i++ {
		temp = temp.Next
	}

	return temp

}

func (L *LL) set(n int, node *Node) *LL {
	if L.Length == 0 {
		return nil
	}
	if n >= L.Length || n < 0 {
		return nil
	}
	temp := L.get(n)
	temp.Value = node.Value

	return L
}

func main() {

	myLL := &LL{}

	// Push into Linked List
	node1 := &Node{Value: 5}
	node2 := &Node{Value: 8}
	node3 := &Node{Value: 1}
	node4 := &Node{Value: 10}

	myLL.push(node1)

	myLL.push(node2)

	// Inspect Linked List
	fmt.Println(myLL.Head)
	fmt.Println(myLL.Tail)

	myLL.push(node3)
	myLL.push(node4)

	test := myLL.get((2))

	fmt.Println(test, "::", myLL.Length)
	fmt.Println(myLL.pop(), "::", myLL.Length)

	test = myLL.get((2))
	fmt.Println(test, "::", myLL.Length)

	myLL.set(1, &Node{Value: 25})

	fmt.Println(myLL.get(1))
}
