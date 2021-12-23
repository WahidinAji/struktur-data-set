package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	set := NewAStringSet()
	assert.True(t, set.Add("Aji"),"Return should be true")
	assert.False(t, set.Add("Aji"),"Return should be False")
	assert.True(t, set.Add("Wahidin"),"Return should be true")
	assert.False(t, set.Add("Wahidin"),"Return should be False")
	assert.True(t, set.Add("Aja"),"Return should be true")
	assert.False(t, set.Add("Aja"),"Return should be False")
	assert.True(t, set.Add("Hehe"),"Return should be true")
	assert.False(t, set.Add("Hehe"),"Return should be False")
	/*
	//assert.Equal(t, set.Add("Aji"),"Aji","They should be equal")
	//assert.Equal(t, set.Add("Wahidin"),"Wahidin","They should be equal")
	//assert.Equal(t, set.Add("Aja"),"Aja","They should be equal")
	//assert.Equal(t, set.Add("Hehe"),"Hehe","They should be equal")
	//
	//assert.NotEqual(t, set.Add("Aji"),"Wahidin","They should be equal")
	//assert.NotEqual(t, set.Add("Wahidin"),"Aji","They should be equal")
	//assert.NotEqual(t, set.Add("Aja"),"Hehe","They should be equal")
	//assert.NotEqual(t, set.Add("Hehe"),"Aja","They should be equal")
	 */
}

func TestContains(t *testing.T) {
	set := NewAStringSet()
	set.Add("Aji")
	set.Add("Wahidin")

	assert.True(t, set.Contains("Aji"))
	assert.True(t, set.Contains("Wahidin"))
	assert.False(t, set.Contains("Aja"))
}

func TestSize(t *testing.T) {
	set := NewAStringSet()
	assert.Equal(t, 0, set.Size(),"Size should be 0")

	set.Add("Aji")
	assert.Equal(t, 1, set.Size(),"Size should be 1")
	set.Add("Aji")
	assert.Equal(t, 1, set.Size(),"Size should be 1")

	set.Add("Wahidin")
	assert.Equal(t, 2, set.Size(),"Size should be 2")
}