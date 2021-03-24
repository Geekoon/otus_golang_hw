package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int			// длина списка
	Front() *ListItem	// первый элемент списка
	Back() *ListItem	// последний элемент списка
	PushFront(v interface{}) *ListItem	// добавить значение в начало
	PushBack(v interface{}) *ListItem	// добавить значение в конец
	Remove(i *ListItem)			// удалить элемент
	MoveToFront(i *ListItem)	// переместить элемент в начало
}

// элемент списка
type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	firstEl *ListItem
	lastEl *ListItem
	len  int
}

// возвращает структуру, удовлетворяющую интерфейсу List
func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

// первый элемент списка
func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.firstEl
}

// последний элемент списка
func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.lastEl
}

// добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	var newEl = ListItem{Value: v}
	if l.len != 0 {
		newEl.Next = l.firstEl
		l.firstEl.Prev = &newEl
	} else {
		l.lastEl = &newEl
	}
	l.firstEl = &newEl
	l.len++

	return &newEl
}

// добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	var newEl = ListItem{Value: v}
	if l.len != 0 {
		newEl.Prev = l.lastEl
		l.lastEl.Next = &newEl
	} else {
		l.firstEl = &newEl
	}
	l.lastEl = &newEl
	l.len++

	return &newEl
}

// Удалить гарантированно существующий элемент
func (l *list) Remove(i *ListItem) {
	if i == l.firstEl {
		l.firstEl = i.Next
		l.firstEl.Prev = nil
	} else
	if i == l.lastEl {
		l.lastEl = i.Prev
		l.lastEl.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
	i.Prev = nil
	i.Next = nil
	i.Value = nil
}

// Переместить элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	if i == l.firstEl {
		return
	}
	if i == l.lastEl {
		l.lastEl = i.Prev
		l.lastEl.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.firstEl.Prev = i
	i.Next = l.firstEl
	i.Prev = nil
	l.firstEl = i
}