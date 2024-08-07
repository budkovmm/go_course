package main

// Deque двунаправленная очередь
type Deque struct {
	elements []int
}

// NewDeque конструктор создающий двунаправленную очередь
func NewDeque(elements ...int) *Deque {
	return &Deque{elements: elements}
}

// PushFront добавляет элемент в начало очереди
func (d *Deque) PushFront(element int) {
	d.elements = append([]int{element}, d.elements...)
}

// PopFront извлекает элемент из начала очереди
func (d *Deque) PopFront() int {
	if len(d.elements) == 0 {
		return 0
	}
	element := d.elements[0]
	d.elements = d.elements[1:]
	return element
}

// PushBack добавляет элемент в конец очереди
func (d *Deque) PushBack(element int) {
	d.elements = append(d.elements, element)
}

// PopBack извлекает элемент из конца очереди
func (d *Deque) PopBack() int {
	l := len(d.elements)

	if l == 0 {
		return 0
	}
	element := d.elements[l-1]
	d.elements = d.elements[:l-1]
	return element
}

func main() {
	deque := NewDeque(1, 2, 3)
	deque.PushFront(0)
	deque.PopFront()
	deque.PopFront()
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushBack(4)
	deque.PopBack()
}
