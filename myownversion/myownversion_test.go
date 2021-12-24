package myownversion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Add Test
func TestAdd(t *testing.T) {
	deps := DependencyStructureDataSet{
		Array: [10]string{},
	}
	assert.True(t, deps.Add("Aji"),"Return should be true")
	assert.False(t, deps.Add("Aji"),"Return should be False")
	assert.True(t, deps.Add("Wahidin"),"Return should be true")
	assert.False(t, deps.Add("Wahidin"),"Return should be False")
	assert.True(t, deps.Add("Aja"),"Return should be true")
	assert.False(t, deps.Add("Aja"),"Return should be False")
	assert.True(t, deps.Add("Hehe"),"Return should be true")
	assert.False(t, deps.Add("Hehe"),"Return should be False")
	fmt.Println(deps.Size(), len(deps.Array))
}

//Contains Test
func TestContains(t *testing.T) {
	deps := DependencyStructureDataSet{}
	deps.Add("Aji")
	deps.Add("Wahidin")
	fmt.Println(deps.Contains("Aji"))
	fmt.Println(deps.Contains("Wahidin"))


	assert.True(t, deps.Contains("Aji"))
	assert.True(t, deps.Contains("Wahidin"))
	assert.False(t, deps.Contains("Aja"))
}

//Size Test
func TestSize(t *testing.T) {
	deps := DependencyStructureDataSet{}
	assert.Equal(t, 0, deps.Size(),"Size should be 0")

	deps.Add("Aji")
	assert.Equal(t, 1, deps.Size(),"Size should be 1")

	deps.Add("Wahidin")
	assert.Equal(t, 2, deps.Size(),"Size should be 2")
}

//Remove Test
func TestRemove(t *testing.T) {
	deps := DependencyStructureDataSet{}
	deps.Add("Aji")
	deps.Add("Wahidin")
	deps.Add("Aja")
	deps.Add("Budi")
	deps.Add("Nugraha")
	deps.Add("Joko")

	assert.Equal(t, 6,deps.Size(),"Size should be <- 6")

	deps.Remove("Wahidin")
	assert.Equal(t, 5,deps.Size(),"Size should be <- 5")

	assert.True(t, deps.Contains("Aji"),"Aji should be <- True")
	assert.True(t, deps.Contains("Aja"),"Aja should be <- True")
	assert.True(t, deps.Contains("Budi"),"Budi should be <- True")
	assert.True(t, deps.Contains("Nugraha"),"Nugraha should be <- True")
	assert.True(t, deps.Contains("Joko"),"Joko should be <- True")
	assert.False(t, deps.Contains("Wahidin"),"Wahidin should be <- True")
}

//Test Growth With Dynamic Array/SLice
func TestGrowth(t *testing.T)  {
	deps := DependencyStructureDataSet{}
	for i := 0; i < 11; i++ {
		deps.Added("Value to" + fmt.Sprint(i),11)
		fmt.Println(deps.ExampleArray)
	}
	fmt.Println(deps.Size())

	assert.Equal(t, len(deps.ExampleArray), deps.Size(),"Error bcz the size from Object of Array is constant <- "+fmt.Sprint(len(deps.ExampleArray)))
}

//Test Added
func TestAdded(t *testing.T)  {
	d := DependencyStructureDataSet{}
	assert.True(t, d.Added("Aji",11),"True")
	fmt.Println(d.Contains("Aji"))
	fmt.Println(d.Size(), len(d.ExampleArray))
}