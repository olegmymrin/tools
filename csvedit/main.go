package main

import (
	"encoding/csv"
	"flag"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/olegmymrin/tools/errorx"
)

var (
	In   string
	Out  string
	Rows string
)

func init() {
	flag.StringVar(&In, "in", "", "In")
	flag.StringVar(&Out, "out", "", "Out")
	flag.StringVar(&Rows, "rows", "", "Row numbers")
}

func main() {
	flag.Parse()

	indexes := map[int]bool{}
	for _, r := range strings.Split(Rows, ",") {
		idx, err := strconv.Atoi(r)
		errorx.FatalOnErr(err, Rows)
		if idx <= 0 {
			flag.Usage()
			return
		}
		indexes[idx-1] = true
	}

	inFile, err := os.Open(In)
	errorx.FatalOnErr(err)
	defer func() {
		errorx.FatalOnErr(inFile.Close())
	}()
	inReader := csv.NewReader(inFile)

	outFile, err := os.OpenFile(Out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	errorx.FatalOnErr(err)
	defer func() {
		errorx.FatalOnErr(outFile.Close())
	}()
	outWriter := csv.NewWriter(outFile)

	for {
		inRow, err := inReader.Read()
		if err == io.EOF {
			break
		}
		errorx.FatalOnErr(err)

		outRow := make([]string, 0, len(indexes))
		for i, cell := range inRow {
			if indexes[i] {
				outRow = append(outRow, cell)
			}
		}

		errorx.FatalOnErr(outWriter.Write(outRow))
	}
}
