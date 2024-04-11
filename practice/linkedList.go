package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
    head *Node
}

// Insert inserts a new node with the given data at the end of the list
func (list *LinkedList) Insert(data int) {
    newNode := &Node{data: data}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// RemoveByValue removes the first occurrence of the node with the given value
func (list *LinkedList) RemoveByValue(value int) {
    if list.head == nil {
        return
    }

    if list.head.data == value {
        list.head = list.head.next
        return
    }

    prev := list.head
    current := list.head.next
    for current != nil {
        if current.data == value {
            prev.next = current.next
            return
        }
        prev = current
        current = current.next
    }
}

// RemoveAt removes the node at the given position (0-based index)
func (list *LinkedList) RemoveAt(position int) {
    if list.head == nil {
        return
    }

    if position == 0 {
        list.head = list.head.next
        return
    }

    prev := list.head
    current := list.head.next
    for i := 1; current != nil; i++ {
        if i == position {
            prev.next = current.next
            return
        }
        prev = current
        current = current.next
    }
}

// Display prints all the elements in the linked list
func (list *LinkedList) Display() {
    current := list.head
    for current != nil {
        fmt.Print(current.data, " ")
        current = current.next
    }
    fmt.Println()
}

func LinkedListfunc() {
    // Create a new linked list
    list := LinkedList{}

    // Insert elements into the linked list
    list.Insert(1)
    list.Insert(2)
    list.Insert(3)
    list.Insert(4)

    // Display the elements of the linked list
    fmt.Println("Linked List:")
    list.Display()

    // Remove an element by value
    list.RemoveByValue(2)
    fmt.Println("After removing 2:")
    list.Display()

    // Remove an element at a specific position
    list.RemoveAt(1) // Remove element at position 1 (0-based index)
    fmt.Println("After removing element at position 1:")
    list.Display()
}
