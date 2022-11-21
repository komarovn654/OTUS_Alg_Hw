package hw04arrays

import (
	"fmt"
)

type MatrixArray struct {
	domains FactorArray
	len     int // current length
	maxLen  int // domain's max length
}

type itemPos struct {
	domain int
	index  int
}

func (ma *MatrixArray) Size() int {
	return ma.len
}

func (ma *MatrixArray) Get(index int) Item {
	pos := ma.getPosition(index)

	if !ma.isDomainExist(pos.domain) {
		panic(fmt.Errorf("domain %v doesnt exist", pos.domain))
	}

	domain := ma.getDomain(pos.domain)
	return domain.Get(pos.index)
}

func (ma *MatrixArray) Add(t Item) {
	pos := itemPos{
		domain: ma.len / ma.maxLen,
		index:  0, // index does not matter; added to the end of domain
	}

	if !ma.isDomainExist(pos.domain) {
		ma.domains.Add(InitFactorArray(0))
	}

	ma.addToDomain(pos.domain, t)
	ma.len++
}

func (ma *MatrixArray) Insert(t Item, index int) {
	pos := itemPos{
		domain: index / ma.maxLen,
		index:  index % ma.maxLen,
	}

	lastItem := ma.insertInDomain(t, pos.domain, pos.index)
	if ma.isDomainFull(pos.domain) {
		for i := 1; i <= ma.len/ma.maxLen-pos.domain; i++ {
			firstItem := lastItem
			if !ma.isDomainExist(pos.domain + i) {
				ma.Add(firstItem)
				return // if domain doest exist, its last domain
			}
			lastItem = ma.insertInDomain(firstItem, pos.domain+i, 0)
		}
	}

	ma.len++
}

func (ma *MatrixArray) Remove(index int) Item {
	pos := itemPos{
		domain: index / ma.maxLen,
		index:  index % ma.maxLen,
	}

	result := ma.removeFromDomain(pos.domain, pos.index)
	ma.len--

	removed := Item(nil)
	for i := ma.len / ma.maxLen; i > pos.domain; i-- {
		lastItem := removed
		removed = ma.removeFromDomain(i, 0)
		if lastItem != nil {
			ma.addToDomain(i, lastItem)
		}
	}

	if removed != nil {
		ma.addToDomain(pos.domain, removed)
	}
	return result
}

func (ma *MatrixArray) removeFromDomain(domainNum int, index int) Item {
	d := ma.getDomain(domainNum)
	removed := d.Remove(index)
	ma.domains.Set(d, domainNum)

	return removed
}

func (ma *MatrixArray) insertInDomain(t Item, domainNum int, index int) Item {

	d := ma.getDomain(domainNum)
	lastItem := d.Get(d.Size() - 1)
	d.Insert(t, index)

	if !ma.isDomainFull(domainNum) {
		ma.domains.Set(d, domainNum) // change lenght in domain
	}
	return lastItem
}

func (ma *MatrixArray) getDomain(domainNum int) FactorArray {
	domain := ma.domains.Get(domainNum)
	if d, ok := domain.(FactorArray); ok {
		return d
	}
	panic(fmt.Errorf("it is impossible to cast Item to Factor array"))
}

func (ma *MatrixArray) getPosition(index int) itemPos {
	return itemPos{
		domain: index / ma.maxLen,
		index:  index % ma.maxLen,
	}
}

func (ma *MatrixArray) addToDomain(domainNum int, t Item) {
	domain := ma.getDomain(domainNum)
	domain.Add(t)

	ma.domains.Set(domain, domainNum)
}

func (ma *MatrixArray) isDomainFull(domainNum int) bool {
	domain := ma.getDomain(domainNum)

	return domain.Size() >= ma.maxLen
}

func (ma *MatrixArray) isDomainExist(domainNum int) bool {
	if ma.domains.arr == nil {
		return false // array doesnt exist
	}

	if ((ma.len - 1) / ma.maxLen) < domainNum {
		return false
	}

	if ma.domains.Get(domainNum) == nil {
		return false
	}
	return true
}

func InitMatrixArray(lenght int) MatrixArray {
	return MatrixArray{domains: FactorArray{}, len: 0, maxLen: 100}
}
