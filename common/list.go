package common

type List struct {
	// Using an empty struct{} has advantage that it doesn't require any additional space
	// Go's internal map type is optimized for that kind of values
	data map[string]struct{}
}

func NewList() *List {

	return &List{data: make(map[string]struct{})}
}

func (l *List) AddItems(items []string) *List {

	for _, curr := range items {
		l.Add(curr)
	}

	return l
}

func (l *List) Add(item string) *List {

	l.data[item] = struct{}{}

	return l
}

func (l List) Contains(item string) bool {

	_, ret := l.data[item]

	return ret
}

func (l List) Size() int {

	return len(l.data)
}

func (l List) Items() []string {

	if len(l.data) == 0 {
		return []string{}
	}

	ret := make([]string, 0, len(l.data))
	for i := range l.data {
		ret = append(ret, i)
	}

	return ret
}

// IsSimilar returns true if both arrays contains same items
func (l List) IsSimilar(list []string) bool {

	ret := true
	if l.Size() != len(list) {
		ret = false
	} else {
		for _, currItem := range list {
			if !l.Contains(currItem) {
				ret = false
				break
			}
		}
	}

	return ret
}
