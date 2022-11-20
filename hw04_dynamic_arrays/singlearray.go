package hw04arrays

type Item interface{}

type SingleArray struct {
	arr []Item
	len int
}

func (sa *SingleArray) Size() int {
	return sa.len
}

func (sa *SingleArray) Set(t Item, index int) Item {
	sa.arr[index] = t
	return t
}

func (sa *SingleArray) Get(index int) Item {
	return sa.arr[index]
}

func (sa *SingleArray) Add(t Item) {
	newArray := InitSingleArray(sa.len + 1)

	for i := 0; i < sa.len; i++ {
		newArray.arr[i] = sa.arr[i]
	}
	newArray.arr[newArray.len-1] = t

	sa.len = newArray.len
	sa.arr = newArray.arr
}

func (sa *SingleArray) Insert(t Item, index int) {
	newArray := InitSingleArray(sa.len + 1)

	insert := false
	for i := 0; i < sa.len; i++ {
		if i == index {
			newArray.arr[i] = t
			insert = true
		}
		if insert {
			newArray.arr[i+1] = sa.arr[i]
			continue
		}
		newArray.arr[i] = sa.arr[i]
	}

	sa.len = newArray.len
	sa.arr = newArray.arr
}

func (sa *SingleArray) Remove(index int) Item {
	newArray := InitSingleArray(sa.len - 1)

	var removed Item
	remove := false
	for i := 0; i < sa.len; i++ {
		if i == index {
			remove = true
			removed = sa.arr[i]
			continue
		}
		if remove {
			newArray.arr[i-1] = sa.arr[i]
			continue
		}
		newArray.arr[i] = sa.arr[i]
	}

	sa.len = newArray.len
	sa.arr = newArray.arr

	return removed
}

func InitSingleArray(length int) SingleArray {
	return SingleArray{arr: make([]Item, length), len: length}
}
