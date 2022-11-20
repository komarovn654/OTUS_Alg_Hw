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

	newArray := InitFactorArray(fa.len + 1)

	for i := 0; i < fa.len; i++ {
		newArray.arr[i] = fa.arr[i]
	}
	newArray.arr[newArray.len-1] = t

	fa.len = newArray.len
	fa.cap = newArray.cap
	fa.arr = newArray.arr
}

func (fa *FactorArray) Insert(t Item, index int) {
	if fa.cap > fa.len {
		for i := fa.len; i > index; i-- {
			fa.arr[i] = fa.arr[i-1]
		}
		fa.arr[index] = t
		fa.len++
		return
	}

	newArray := InitFactorArray(fa.len + 1)
	insert := false
	for i := 0; i < fa.len; i++ {
		if i == index {
			newArray.arr[i] = t
			insert = true
		}
		if insert {
			newArray.arr[i+1] = fa.arr[i]
			continue
		}
		newArray.arr[i] = fa.arr[i]
	}

	fa.len = newArray.len
	fa.cap = newArray.cap
	fa.arr = newArray.arr
}

func (fa *FactorArray) Remove(index int) Item {
	removed := fa.arr[index]
	for i := index; i < fa.len-1; i++ {
		fa.arr[i] = fa.arr[i+1]
	}
	fa.len--
	return removed
}

func InitFactorArray(length int) FactorArray {
	return FactorArray{arr: make([]Item, length*2), len: length, cap: length * 2}
}
