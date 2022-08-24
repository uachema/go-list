package list

func NewList() List {
	return List{}
}

type List struct {
	data []any
}

func (l *List) All() []any {
	return l.data
}

func (l *List) Update(list []any) {
	l.data = list
}

func (l *List) Add(values ...any) {
	l.data = append(l.data, values...)
}

func (l *List) RemoveByIndex(indexList ...int) {
	for _, v := range indexList {
		l.data = append(l.data[:v], l.data[v+1:]...)
	}
}

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

func (l *List) IndexOf(value any) int {
	for i, v := range l.data {
		if value == v {
			return i
		}
	}
	return -1
}

// Returns nil if value do not exists
func (l *List) OnIndex(index int) any {
	for i, v := range l.data {
		if i == index {
			return v
		}
	}
	return nil
}

func (l *List) Contains(value any) bool {
	for _, v := range l.data {
		if value == v {
			return true
		}
	}
	return false
}

func (l *List) Length() int {
	return len(l.data)
}

func (l *List) Clear() {
	var empty []any
	l.data = empty
}

func BinarySearch(arr []int, x int, low int, hight int) int {
	if hight >= low {
		mid := low + (hight-low)/2

		// If the element is present at the middle
		// itself
		if arr[mid] == x {
			return mid
		}

		// If element is smaller than mid, then
		// it can only be present in left subarray
		if arr[mid] > x {
			return BinarySearch(arr, x, low, mid-1)
		}

		// Else the element can only be present
		// in right subarray
		return BinarySearch(arr, x, mid+1, hight)
	}

	// We reach here when element is not
	// present in array
	return -1
}
