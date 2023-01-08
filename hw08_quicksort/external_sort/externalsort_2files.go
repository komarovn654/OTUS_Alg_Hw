package externalsort

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
)

var (
	ErrEmptyString = errors.New("empty string")
)

func (f *File) ExternalSort2Files() (err error) {
	f.sub1.path = "subfile1.txt"
	f.sub2.path = "subfile2.txt"

	f.file, err = os.OpenFile(f.path, os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.file.Close()

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

	return f.externalSort2Files(1)
}

func (f *File) externalSort2Files(sortedLen int) error {
	if sortedLen > f.lines {
		return nil
	}

	if err := f.split2Files(sortedLen); err != nil {
		return err
	}
	if err := f.merge2Files(sortedLen); err != nil {
		return err
	}

	return f.externalSort2Files(sortedLen * 2)
}

func (f *File) split2Files(sortedLen int) error {
	if err := f.eraseSubFiles(); err != nil {
		return err
	}
	if _, err := f.file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	source := bufio.NewScanner(f.file)
	for i := 0; source.Scan(); i++ {
		if i > sortedLen*2-1 {
			i = 0
		}
		if i < sortedLen {
			if _, err := f.sub1.file.WriteString(source.Text() + "\n"); err != nil {
				return err
			}
			continue
		}
		if _, err := f.sub2.file.WriteString(source.Text() + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (f *File) merge2Files(sortedLen int) error {
	if err := f.eraseSourceFile(); err != nil {
		return err
	}
	if err := f.seekSubFiles(0); err != nil {
		return err
	}

	sub1 := bufio.NewScanner(f.sub1.file)
	sub2 := bufio.NewScanner(f.sub2.file)

	v1 := f.sub1.getItem(sub1)
	v2 := f.sub2.getItem(sub2)
	for !f.sub1.isEnd || !f.sub2.isEnd {
		f.sub1.sortLen = sortedLen
		f.sub2.sortLen = sortedLen
		for i := 0; (i < sortedLen*2) && (!f.sub1.isEnd || !f.sub2.isEnd); i++ {
			if ((v1 < v2) && f.sub1.sortLen > 0) || f.sub2.sortLen == 0 || f.sub2.isEnd {
				f.file.WriteString(strconv.Itoa(v1) + "\n")
				v1 = f.sub1.getItem(sub1)
				f.sub1.sortLen--
				continue
			}
			f.file.WriteString(strconv.Itoa(v2) + "\n")
			v2 = f.sub2.getItem(sub2)
			f.sub2.sortLen--
		}
	}

	return nil
}
