package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tufin/totem/analysis"
)

func TestArray_AddItems(t *testing.T) {

	arr := analysis.NewList().AddItems([]string{"a", "c"}).Add("b")
	assert.True(t, arr.IsSimilar([]string{"a", "b", "c"}))
	assert.Equal(t, 3, arr.Size())
}
func TestArray_Contains(t *testing.T) {

	arr := analysis.NewList().AddItems([]string{"b", "a"})
	assert.True(t, arr.Contains("a"))
	assert.True(t, arr.Contains("b"))
	assert.False(t, arr.Contains("c"))
}

func TestArray_IsSimilar_True(t *testing.T) {

	arr := analysis.NewList()
	arr.AddItems([]string{"a", "b", "c"})
	assert.True(t, arr.IsSimilar([]string{"b", "c", "a"}))
}

func TestArray_Size(t *testing.T) {

	arr := analysis.NewList()
	arr.AddItems([]string{"a", "b", "c"})
	assert.Equal(t, 3, arr.Size())
}

func TestArray_Items(t *testing.T) {

	const name = "nehmad"
	list := analysis.NewList().Add(name)
	assert.Equal(t, 1, len(list.Items()))
	assert.True(t, list.Contains(name))
}

func TestArray_ItemsEmpty(t *testing.T) {

	assert.Equal(t, 0, len(analysis.NewList().Items()))
}

func TestArray_IsSimilarEmptyArrays(t *testing.T) {

	arr := analysis.NewList()
	assert.True(t, arr.IsSimilar([]string{}))
}

func TestArray_IsSimilarSize_False(t *testing.T) {

	arr := analysis.NewList()
	arr.AddItems([]string{"a", "c"})
	assert.False(t, arr.IsSimilar([]string{"b", "c", "a"}))
}

func TestArray_IsSimilar_False(t *testing.T) {

	arr := analysis.NewList()
	arr.AddItems([]string{"a", "c", "d"})
	assert.False(t, arr.IsSimilar([]string{"b", "c", "a"}))
}
