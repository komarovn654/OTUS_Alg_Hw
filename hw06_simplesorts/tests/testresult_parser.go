package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	hw06 "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
)

type Rows map[string][]hw06.SortTime
type ResultTable map[string]Rows

func fillResultTable(r TestResult, rt *ResultTable) *ResultTable {
	if (*rt)[r.ArrayType] == nil {
		(*rt)[r.ArrayType] = make(Rows)
	}
	if (*rt)[r.ArrayType][r.SortName] == nil {
		(*rt)[r.ArrayType][r.SortName] = make([]hw06.SortTime, 8)
	}

	(*rt)[r.ArrayType][r.SortName][int(math.Log10(float64(r.N)))] = r.SortTime
	// switch r.N {
	// case 1:
	// 	(*rt)[r.ArrayType][r.SortName][int(math.Log10(float64(r.N)))] = r.SortTime
	// case 10:
	// 	(*rt)[r.ArrayType][r.SortName][1] = r.SortTime
	// case 100:
	// 	(*rt)[r.ArrayType][r.SortName][2] = r.SortTime
	// case 1000:
	// 	(*rt)[r.ArrayType][r.SortName][3] = r.SortTime
	// case 10000:
	// 	(*rt)[r.ArrayType][r.SortName][4] = r.SortTime
	// case 100000:
	// 	(*rt)[r.ArrayType][r.SortName][5] = r.SortTime
	// case 1000000:
	// 	(*rt)[r.ArrayType][r.SortName][6] = r.SortTime
	// case 10000000:
	// 	(*rt)[r.ArrayType][r.SortName][7] = r.SortTime
	// default:
	// 	return nil
	// }

	return rt
}

func saveResultTable(filePath string, rt *ResultTable) error {
	f, err := os.OpenFile(filePath, os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	var sb strings.Builder
	for table := range *rt {
		sb.WriteString(fmt.Sprintf("# %v\n", table))
		sb.WriteString("| Sort name | 1 | 10 | 100 | 1000 | 10000 | 100000 | 1000000 | 10000000 |\n")
		sb.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
		for rowName, row := range (*rt)[table] {
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