package hw04arrays

import "fmt"

type DynamicArray struct {
	arr []Item
	len int
	cap int
}

func (da *DynamicArray) Size() int {
	return da.len
}

func (da *DynamicArray) Set(t Item, index int) Item {
	if index > da.len {
		panic(fmt.Errorf("index %v out of range %v", index, da.len))
	}

	da.arr[index] = t
	return t
}

func (da *DynamicArray) Get(index int) Item {
	if index > da.len {
		panic(fmt.Errorf("index %v out of range %v", index, da.len))
	}

	return da.arr[index]
}

func (da *DynamicArray) Add(t Item) {
	if da.cap > da.len {
		da.arr[da.len] = t
		da.len++
		return
	}

	newArray := InitDynamicArray(da.len + 1)

	for i := 0; i < da.len; i++ {
		newArray.arr[i] = da.arr[i]
	}
	newArray.arr[newArray.len-1] = t

	da.len = newArray.len
	da.cap = newArray.cap
	da.arr = newArray.arr
}

func (da *DynamicArray) Insert(t Item, index int) {
	if da.cap > da.len {
		for i := da.len; i > index; i-- {
			da.arr[i] = da.arr[i-1]
		}
		da.arr[index] = t
		da.len++
		return
	}

	newArray := InitDynamicArray(da.len + 1)
	insert := false
	for i := 0; i < da.len; i++ {
		if i == index {
			newArray.arr[i] = t
			insert = true
		}
		if insert {
			newArray.arr[i+1] = da.arr[i]
			continue
		}
		newArray.arr[i] = da.arr[i]
	}

	da.len = newArray.len
	da.cap = newArray.cap
	da.arr = newArray.arr
}

func (da *DynamicArray) Remove(index int) Item {
	removed := da.arr[index]
	for i := index; i < da.len-1; i++ {
		da.arr[i] = da.arr[i+1]
	}
	da.len--
	return removed
}

func InitDynamicArray(length int) DynamicArray {
	return DynamicArray{arr: make([]Item, length+5), len: length, cap: length + 5}
}
