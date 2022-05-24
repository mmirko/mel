package mel

func orderedPlace(queue_head **individual, newindiv *individual) {
	head := *queue_head
	tail := head
	if head == nil {
		*queue_head = newindiv
	} else {
		for curr := head; curr != nil; curr = curr.next {
			tail = curr
			if newindiv.fitness_values[0] > curr.fitness_values[0] {
				newindiv.prev = curr.prev
				newindiv.next = curr
				if curr.prev != nil {
					curr.prev.next = newindiv
				} else {
					*queue_head = newindiv
				}
				curr.prev = newindiv
				return
			}
		}
		newindiv.prev = tail
		tail.next = newindiv

	}
}
