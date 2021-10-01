package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	followingNode := l.head
	l.head = n
	l.head.next = followingNode
	l.length++
}

func (l linkedList) printListValues() {
	nodeToPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", nodeToPrint.data)
		nodeToPrint = nodeToPrint.next
		l.length--
	}
}

func main() {
	linkedList := linkedList{head: &node{}}
	node1 := node{
		data: 23,
		next: &node{},
	}
	node2 := node{
		data: 3,
		next: &node{},
	}
	linkedList.prepend(&node1)
	linkedList.prepend(&node2)
	fmt.Println(linkedList)
	linkedList.printListValues()
}
