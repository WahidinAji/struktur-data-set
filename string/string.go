package string

import "fmt"

type StringSet struct {
	Set map[string]bool
}

func NewAStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(i string) bool {
	_, found := set.Set[i]
	set.Set[i] = true
	return !found //False if it existed already
}

func (set *StringSet) Contains(i string) bool {
	_, found := set.Set[i]
	return found //true if it existed already
}

func (set *StringSet) Remove(i string) {
	delete(set.Set, i)
}


func (set *StringSet) Size() int {
	return len(set.Set)
}

func Start()  {
	set := NewAStringSet()
	set.Add("Aji")
	set.Add("Wahidin")
	set.Add("Aja")
	set.Add("Hehe")
	set.Add("Hehe")
	fmt.Println(set.Size())
	fmt.Println(set.Contains("Wahidin"))
	set.Remove("Wahidin")
	fmt.Println(set.Size())
	fmt.Println(set.Contains("Wahidin"))
	fmt.Println(set.Contains("Hehe"))
}

