package externalsort

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"

	hw08quicksort "github.com/komarovn654/OTUS_Alg_Hw/hw08_quicksort"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

const (
	SORT_LEN = 100
)

func (f *File) ExternalSortPresort() (err error) {
	f.sub1.path = "subfile1.txt"
	f.sub2.path = "subfile2.txt"

	f.file, err = os.OpenFile(f.path, os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.file.Close()

	if f.lines < SORT_LEN {
		return f.sortWithStash()
	}

	f.sub1.file, err = os.OpenFile(f.sub1.path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.sub1.file.Close()

	f.sub2.file, err = os.OpenFile(f.sub2.path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.sub2.file.Close()

	if err := f.presortSplit(SORT_LEN); err != nil {
		return err
	}

	return f.externalSortPresort(SORT_LEN)
}

func (f *File) externalSortPresort(sortedLen int) error {
	if sortedLen > f.lines {
		return nil
	}

	if err := f.merge2Files(sortedLen); err != nil {
		return err
	}
	if err := f.split2Files(sortedLen * 2); err != nil {
		return err
	}

	return f.externalSortPresort(sortedLen * 2)
}

func (f *File) presortSplit(sortLen int) error {
	source := bufio.NewScanner(f.file)
	var ar []sortutils.Item
	var err error
	isEOF := false

	i := 0
	for len := 0; len < f.lines && !isEOF; len = len + sortLen {
		ar, isEOF, err = f.readArray(source, sortLen)
		if err != nil {
			return err
		}

		if i > sortLen*2-1 {
			i = 0
		}
		if i < sortLen {
			if _, err := f.sub1.file.WriteString(sortArray(ar)); err != nil {
				return err
			}
		} else {
			if _, err := f.sub2.file.WriteString(sortArray(ar)); err != nil {
				return err
			}
		}
		i = i + 100
	}
	return nil
}

func (f *File) sortWithStash() error {
	source := bufio.NewScanner(f.file)
	ar, _, err := f.readArray(source, f.lines)
	if err != nil {
		return err
	}

	if err := f.eraseSourceFile(); err != nil {
		return err
	}

	if _, err := f.file.WriteString(sortArray(ar)); err != nil {
		return err
	}

	return nil
}

func (f *File) readArray(source *bufio.Scanner, len int) ([]sortutils.Item, bool, error) {
	array := make([]sortutils.Item, 0)
	i := 0
	for ; (i < len) && source.Scan(); i++ {
		v, err := strconv.Atoi(source.Text())
		if err != nil {
			return nil, false, err
		}
		array = append(array, sortutils.Item(v))
	}

	return array, i < len, nil
}

func sortArray(array []sortutils.Item) string {
	var strArray strings.Builder
	a := sortutils.Array{Ar: array}
	ctx := context.Background()

	a.SortArray(ctx, hw08quicksort.MergeSort)

	for _, v := range a.Ar {
		strArray.WriteString(strconv.FormatInt(int64(v), 10) + "\n")
	}

	return strArray.String()
}
