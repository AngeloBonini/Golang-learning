package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	length int
	head   *node
	tail   *node
}

func (size *list) Leng() int {
	return size.length
}

func (linkedList *list) addHead(n *node) {
	if linkedList.head == nil {
		linkedList.head = n
		linkedList.tail = n
		linkedList.length++
	} else {
		linkedList.tail.next = n
		linkedList.tail = n
		linkedList.length++
	}
}
func (l list) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
}

func main() {
	lista := list{}
	node1 := &node{data: 20}
	node2 := &node{data: 30}
	node3 := &node{data: 40}
	node4 := &node{data: 50}
	node5 := &node{data: 70}
	lista.addHead(node1)
	lista.addHead(node2)
	lista.addHead(node3)
	lista.addHead(node4)
	lista.addHead(node5)
	fmt.Println("Length = ", lista.Leng())
	lista.Display()

}
