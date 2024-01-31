package list

// Create a new List
func NewList() List {
	return List{}
}

type List struct {
	data []any
}

// Returns all elements of the List
func (l *List) All() []any {
	return l.data
}

// Updates the list to new given list,
// Removes all the previous data
func (l *List) UpdateList(list []any) {
	l.data = list
}

// Updates the value of specified index,
func (l *List) UpdateOn(index int, value any) {
	l.data[index] = value
}

// Add values to the list
// You can pass multiple values as argument
func (l *List) Add(values ...any) {
	l.data = append(l.data, values...)
}

// Removes Specified Index from the List,
// You can pass multiple arguments
func (l *List) RemoveByIndex(indexList ...int) {
	for _, i := range indexList {
		if i >= 0 {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}
}

// Removes Specified Values from the List,
// You can pass multiple arguments
func (l *List) RemoveByValue(valueList ...any) {
	for _, value := range valueList {
		index := -1
		for i, v := range l.data {
			if value == v {
				index = i
				break
			}
		}
		if index >= 0 {
			l.data = append(l.data[:index], l.data[index+1:]...)
		}
	}
}

// Returns index of given value
func (l *List) IndexOf(value any) int {
	for i, v := range l.data {
		if value == v {
			return i
		}
	}
	return -1
}

// Returns value of given index,
// Returns nil if value do not exists
func (l *List) OnIndex(index int) any {
	for i, v := range l.data {
		if i == index {
			return v
		}
	}
	return nil
}

// Returns true if List contains specified value
func (l *List) Contains(value any) bool {
	for _, v := range l.data {
		if value == v {
			return true
		}
	}
	return false
}

// Return lenght of the List
func (l *List) Length() int {
	return len(l.data)
}

// Deletes all the data from List
func (l *List) Clear() {
	var empty []any
	l.data = empty
}
func (l *List) Clr() {
	var empty []any
	l.data = empty
}
