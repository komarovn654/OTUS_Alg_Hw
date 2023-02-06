package openaddressing

type hashTable struct {
	size     int
	capacity int
}

func InitOAHashTable() hashTable {
	return hashTable{}
}
