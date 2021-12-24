package pzn

//var Array = [10]string{}
//var size = 0

type SetArr struct {
	Set [10]string
	size int
}

func NewASet() *SetArr {
	return &SetArr{Set: [10]string{}, size: 0}
}

func (set *SetArr) Add(value string) bool {
	if set.Contains(value) {
		return false
	} else {
		set.Set[set.size] = value
		set.size++
		return true
	}
}

func (set *SetArr) Contains(value string) bool {
	for _, item := range set.Set {
		if value == item {
			return true
		}
	}
	return false;
}

func (set *SetArr) Size() int {
	return set.size;
}
func Remove(value string) bool {
	return true;
}
