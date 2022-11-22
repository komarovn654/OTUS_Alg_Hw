package hw04arrays

import "fmt"

type FactorArray struct {
	arr []Item
	len int
	cap int
}

func (fa *FactorArray) Size() int {
	return fa.len
}

func (fa *FactorArray) Set(t Item, index int) Item {
	if index > fa.len {
		panic(fmt.Errorf("index %v out of range %v", index, fa.len))
	}

	fa.arr[index] = t
	return t
}

func (fa *FactorArray) Get(index int) Item {
	if index > fa.len {
		panic(fmt.Errorf("index %v out of range %v", index, fa.len))
	}

	return fa.arr[index]
}

func (fa *FactorArray) Add(t Item) {
	if fa.cap > fa.len {
		fa.arr[fa.len] = t
		fa.len++
		return
	}

	fa.resize(fa.len + 1)
	fa.arr[fa.len-1] = t
}

func (fa *FactorArray) Insert(t Item, index int) {
	if fa.cap <= fa.len {
		fa.resize(fa.len + 1)
		fa.len--
	}

	for i := fa.len; i > index; i-- {
		fa.arr[i] = fa.arr[i-1]
	}
	fa.arr[index] = t
	fa.len++
}

func (fa *FactorArray) Remove(index int) Item {
	removed := fa.arr[index]
	for i := index; i < fa.len-1; i++ {
		fa.arr[i] = fa.arr[i+1]
	}
	fa.len--
	return removed
}

func (fa *FactorArray) resize(length int) {
	newArray := InitFactorArray(length)
	for i := 0; i < fa.len; i++ {
		newArray.arr[i] = fa.arr[i]
	}

	fa.arr = newArray.arr
	fa.len = newArray.len
	fa.cap = newArray.cap
}

func InitFactorArray(length int) FactorArray {
	return FactorArray{arr: make([]Item, length*2), len: length, cap: length * 2}
}
