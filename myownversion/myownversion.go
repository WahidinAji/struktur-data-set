package myownversion

import "fmt"

type DependencyStructureDataSet struct {
	Array  [10]string
	Length int
	ExampleArray []string
}

//use this one for custom length and dynamic length of object array/slice
func (d *DependencyStructureDataSet) Added(value string, lengthSlice int) bool {
	d.ExampleArray = make([]string,lengthSlice)
	if d.Contains(value) {
		return false
	} else {
		d.ensure()
		d.ExampleArray[d.Length] = value
		d.Length++
		return true
	}
}
//to ensure the length of array/slice
func (d *DependencyStructureDataSet) ensure() {
	if len(d.ExampleArray) >= 10 {
		var temp = make([]string,len(d.ExampleArray))
		for i := 0; i < len(d.ExampleArray); i++ {
			temp[i] = d.ExampleArray[i]
		}
		d.ExampleArray = temp
	}
}

//this one is using fixed length of array
func (deps *DependencyStructureDataSet) Add(value string) bool {
	if deps.Contains(value) {
		return false
	} else {
		//deps.ensureCapacity(len(deps.Array))
		deps.Array[deps.Length] = value
		deps.Length++
		//fmt.Println(deps.Length)
		return true
	}
}

//func (deps *DependencyStructureDataSet) ensureCapacity(value int)  {
//	if deps.Length >= len(deps.Array)  {
//		//c := make([]string, len(deps.Array) *2)
//		//copy(c,deps.Array[:value])
//		//c[value] = fmt.Sprint(value)
//		//copy(c[value*2:],deps.Array[value:])
//
//		//create new array
//		var temp = make([]string,len(deps.Array) * 2)
//		for i := 0; i < len(deps.Array); i++ {
//			temp[i] = deps.Array[i] //move data from Array to temp
//		}
//		//then put it into object Array again
//		deps.Array = temp
//	}
//}

func (deps *DependencyStructureDataSet) Contains(value string) bool {
	for _, item := range deps.Array {
		if value == item {
			return true
		}
	}
	return false;
}

func (deps *DependencyStructureDataSet) Size() int {
	return deps.Length;
}

func (deps *DependencyStructureDataSet)indexOf(value string) int {
	for i := 0; i < len(deps.Array); i++ {
		if value == deps.Array[i] {
			return i
		}
	}
	return -1
}

func (deps *DependencyStructureDataSet) Remove(value string) bool {
	if deps.Contains(value) {
		//remove Wahidin
		//[Aji, Wahidin, Aja, Budi, Nugraha, Omoy,null,null,null,null]
		//size ==> 6
		var indexRemoved int = deps.indexOf(value)
		for i := indexRemoved; i < deps.Length; i++ {
			deps.Array[i] = deps.Array[i+1]
		}
		deps.Length--
		return true;
	} else {
		return false
	}
}

func Starting()  {
	deps := DependencyStructureDataSet{}
	deps.Add("Aji")
	deps.Add("Wahidin")
	deps.Add("Aja")
	deps.Add("Hehe")
	deps.Add("Hehe")
	fmt.Println("Get length <- : \t ",deps.Size())
	fmt.Println("Get value Wahidin from object `Array` of Struct <- : \t ",deps.Contains("Wahidin"))
	deps.Remove("Wahidin")
	fmt.Println("Get length after remove Wahidin from object 'Array' <- : \t ",deps.Size())
	fmt.Println("Check Wahidin is exist or not <- : \t ",deps.Contains("Wahidin"))
	fmt.Println("Check Hehe is exist or not <- : \t ",deps.Contains("Hehe"))
}