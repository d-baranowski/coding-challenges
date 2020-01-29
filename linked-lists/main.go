package main

import "math"

type Node struct {
	Next *Node
	Data int
}

func (head *Node) print() *Node {
	n := head

	for n.Next != nil {
		print(n.Data)
		n = n.Next
	}
	print(n.Data)

	return head
}

func (head *Node) add(data []int) *Node {
	t := head

	for _, e := range data {
		t.Next = &Node{
			Next: nil,
			Data: e,
		}
		t = t.Next
	}

	return head
}

func (node *Node) delete() bool {
	if node == nil || node.Next == nil {
		return false
	}

	node.Data = node.Next.Data
	node.Next = node.Next.Next
	return true
}

func (head *Node) removeDuplicates() *Node {
	current := head

	for current != nil {
		runner := current

		for runner.Next != nil {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		current = current.Next
	}

	return head
}

func (head *Node) findElementFromLast(k int) *Node {
	current := head
	runner := head
	end := 0

	for runner.Next != nil {
		end++
		runner = runner.Next
	}

	if end < k {
		return nil
	}

	for i := 0; i < end - k; i++ {
		current = current.Next
	}

	return current
}

func (head *Node) deleteMiddleElement() {
	counter := 0
	current := head

	for current != nil {
		counter++
		current = current.Next
	}

	middle := int(math.Round(float64(counter) / 2))

	current = head
	for i := 0; i < middle - 1; i++ {
		current = current.Next
	}

	current.delete()
}

func main() {
	linkedList := Node{
		Next: nil,
		Data: 9,
	}

	data := []int{1, 2, 3, 1, 2, 3, 3, 2, 1, 4, 5, 6, 7, 9, 9, 9, 1}

	linkedList.add(data)
	linkedList.print()

	print("\n")
	//
	//linkedList.removeDuplicates()
	//
	//linkedList.print()
	//print(linkedList.findElementFromLast(1).Data)
	linkedList.deleteMiddleElement()

	linkedList.print()

	print("\n")

}
