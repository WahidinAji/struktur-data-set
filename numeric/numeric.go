package numeric

import "fmt"

type MapIntSet struct {
	set map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{make(map[int]bool)}
}

func (set *MapIntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *MapIntSet) Contains(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *MapIntSet) Remove(i int) {
	delete(set.set, i)
}

func (set *MapIntSet) Size() int {
	return len(set.set)
}

func Start()  {
	set := NewMapIntSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	fmt.Println(set.Size())
	fmt.Println(set.Contains(2))
	set.Remove(2)
	fmt.Println(set.Size())
	fmt.Println(set.Contains(2))
	fmt.Println(set.Contains(3))
	fmt.Println(set.Contains(4))
}
