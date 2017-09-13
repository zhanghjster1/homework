package main

type Item struct {
	next *Item
	value int
}

type List struct {
	sentinel Item
}

func RevertLinkList(list *List) {
	var head *Item
	for cur := list.sentinel.next; cur != nil; {
		// save the next item pointer
		next := cur.next

		// link head after the current item
		cur.next = head

		// move head to the current item
		head = cur

		// move current item to next
		cur = next
	}

	list.sentinel.next = head
}

func main() {
	var list = List{}

	cur := &list.sentinel

	for i :=0; i<100; i++ {
		cur.next = &Item{
			value: i,
		}
		cur = cur.next
	}

	RevertLinkList(&list)

	for head := list.sentinel.next;head != nil; head = head.next {
		println(head.value)
	}
}
