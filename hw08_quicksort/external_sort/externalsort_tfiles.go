package externalsort

import (
	"bufio"
	"os"
	"strconv"
)

func (f *File) ExternalSortTFiles() error {
	tmpDir := ".splited"
	if err := f.splitTFiles(".splited"); err != nil {
		return err
	}

	if err := f.mergeTFiles(tmpDir); err != nil {
		return nil
	}

	return nil
}

func (f *File) splitTFiles(dir string) error {
	if err := os.Mkdir(dir, os.ModeDir); err != nil {
		return err
	}

	file, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		writeItemTFiles(dir, fs.Text())
	}

	return err
}

func (f *File) mergeTFiles(dir string) error {
	// files, err := os.ReadDir(dir)
	// if err != nil {
	// 	return err
	// }
	df, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer df.Close()

	for i := 0; i < f.max; i++ {
		if ok := isFileExist(dir, i); !ok {
			continue
		}

		items, err := os.ReadFile(dir + "/" + strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		if _, err := df.Write(items); err != nil {
			return err
		}
	}

	return err
}

func isFileExist(dir string, num int) bool {
	if _, err := os.Stat(dir + "/" + strconv.Itoa(num) + ".txt"); err == nil {
		return true
	}
	return false
}

func writeItemTFiles(dir string, item string) error {
	if _, err := strconv.Atoi(item); err != nil {
		return err
	}

	file := dir + "/" + item + ".txt"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(item + "\n")
	return err
}
