package hw04arrays

type sliceWrap struct {
	slice []Item
}

func (sw *sliceWrap) Get(index int) Item {
	return sw.slice[index]
}

func (sw *sliceWrap) Add(t Item) {
	sw.slice = append(sw.slice, t)
}

func (sw *sliceWrap) Insert(t Item, index int) {
	if len(sw.slice) == index {
		sw.slice = append(sw.slice, t)
		return
	}
	ns := append(sw.slice[:index+1], sw.slice[index:]...)
	ns[index] = t
	sw.slice = ns
}

func (sw *sliceWrap) Remove(index int) Item {
	if len(sw.slice) == 0 {
		return nil
	}

	v := sw.slice[index]
	if len(sw.slice) == 1 {
		sw.slice = []Item{}
		return v
	}

	sw.slice = append(sw.slice[:index], sw.slice[index+1:]...)
	return v
}

func (sw *sliceWrap) Size() int {
	return len(sw.slice)
}

func InitSliceWrap() sliceWrap {
	return sliceWrap{slice: []Item{}}
}
