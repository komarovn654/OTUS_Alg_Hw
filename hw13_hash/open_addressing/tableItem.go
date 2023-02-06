package openaddressing

type key interface{}
type value interface{}

type tableItem struct {
	k key
	v value
}
