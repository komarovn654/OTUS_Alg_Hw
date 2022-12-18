package sortutils

import (
	"fmt"
	"os"
	"strings"
)

func SaveResultTable(filePath string, rt resultTable) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	var sb strings.Builder
	for table := range rt {
		sb.WriteString(fmt.Sprintf("# %v\n", table))
		sb.WriteString("| Sort name | 1 | 10 | 100 | 1000 | 10000 | 100000 | 1000000 | 10000000 |\n")
		sb.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
		for rowName, row := range rt[table] {
			sb.WriteString(fmt.Sprintf("| %v |", rowName))
			for _, t := range row {
				if t.Timeout {
					sb.WriteString(" timeout |")
					continue
				}
				sb.WriteString(fmt.Sprintf(" %v |", t.Time))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	_, err = f.WriteString(sb.String())
	return err
}
