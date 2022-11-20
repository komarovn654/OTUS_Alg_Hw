package hw04arrays

type MatrixArray struct {
	domains FactorArray
	len     int // current length
	maxLen  int // domain's max length
}

func (ma *MatrixArray) Size() int {
	return ma.len
}

func (ma *MatrixArray) Add(t Item) {
	domainNum := ma.len / ma.maxLen // domain for next item

	if !ma.isDomainExsist(domainNum) {
		ma.domains.Add(InitFactorArray(0))
	}

	ma.addToDomain(domainNum, t)
	ma.len++
}

func (ma *MatrixArray) addToDomain(domainNum int, t Item) {
	domain := ma.domains.Get(domainNum)
	if d, ok := domain.(FactorArray); ok {
		d.Add(t)
		ma.domains.Set(d, domainNum)
		return
	}
}

func (ma *MatrixArray) isDomainExsist(domainNum int) bool {
	if ma.domains.arr == nil {
		return false // array doesnt exsist
	}

	if ((ma.len - 1) / ma.maxLen) < domainNum {
		return false
	}

	if ma.domains.Get(domainNum) == nil {
		return false
	}
	return true
}
