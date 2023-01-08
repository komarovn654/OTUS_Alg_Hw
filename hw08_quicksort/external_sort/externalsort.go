package externalsort

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrTimeout = errors.New("sort timeout")
)

type File struct {
	file  *os.File
	path  string
	lines int
	max   int
	sub1  SubFile
	sub2  SubFile
}

type SubFile struct {
	file    *os.File
	path    string
	sortLen int
	isEnd   bool
}

// func (f *File) SortFile(sortFunction func() error) (time.Duration, error) {
// 	var err error
// 	newCtx, cancel := context.WithTimeout(context.Background(), time.Second*120)
// 	defer cancel()
// 	sTime := make(chan time.Duration)

// 	go func() {
// 		t := time.Now()
// 		err = f.sortFunction()
// 		sTime <- time.Since(t)
// 	}()

// 	select {
// 	case <-newCtx.Done():
// 		return 0, ErrTimeout
// 	case t := <-sTime:
// 		return t, err
// 	}
// }

func InitRandFile(path string, n int, t int) (File, error) {
	file, err := os.Create(path)
	if err != nil {
		return File{}, err
	}
	defer file.Close()

	str, err := randomizer(n, t)
	if err != nil {
		return File{}, err
	}

	_, err = file.WriteString(str)
	if err != nil {
		return File{}, err
	}

	return File{path: path, lines: n, max: t}, nil
}

func (f *File) IsSorted() (bool, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	stash := 0
	fs := bufio.NewScanner(file)
	for fs.Scan() {
		v, err := strconv.Atoi(fs.Text())
		if err != nil {
			return false, err
		}
		if v < stash {
			return false, nil
		}
		stash = v
	}

	return true, nil
}

func (f *File) eraseSubFiles() error {
	if _, err := f.sub1.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := f.sub1.file.Truncate(0); err != nil {
		return err
	}
	if _, err := f.sub2.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	return f.sub2.file.Truncate(0)
}

func (f *File) seekSubFiles(offset int64) error {
	if _, err := f.sub1.file.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	if _, err := f.sub2.file.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	return nil
}

func (f *File) eraseSourceFile() error {
	if _, err := f.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	return f.file.Truncate(0)
}

func (sf *SubFile) getItem(scan *bufio.Scanner) int {
	if sf.isEnd = !scan.Scan(); sf.isEnd {
		return 0
	}
	v, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic("TODO: i shouldnt do it =(")
	}
	return v
}

func randomizer(n int, t int) (string, error) {
	var sb strings.Builder
	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < n-1; i++ {
		if _, err := sb.WriteString(strconv.FormatInt(s1.Int63()%int64(t), 10) + "\n"); err != nil {
			return "", err
		}
	}
	if _, err := sb.WriteString(strconv.FormatInt(s1.Int63()%int64(t), 10)); err != nil {
		return "", err
	}

	return sb.String(), nil
}
