package ds

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (l *LinkedList) Prepend(data int) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
}
func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}
	temp := l.head
	for temp.next != nil && temp.next.data != value {
		temp = temp.next
	}
	if temp.next != nil {
		temp.next = temp.next.next
	}
}

func (l *LinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}
