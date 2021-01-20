package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
)

var (
	In   string
	Out  string
	Rows string
)

func init() {
	flag.StringVar(&In, "in", "", "In")
	flag.StringVar(&Out, "out", "", "Out")
	flag.StringVar(&Out, "out", "", "Out")
}

func main() {
	flag.Parse()
	in, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	recs, err := csv.NewReader(in).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.OpenFile(Out, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var no int
	for i, r := range recs {
		if i == 0 || r[0] != recs[i-1][0] {
			no = 1
		}
		r[1] = r[1] + "-" + strconv.Itoa(no)
		no++
	}
	err = csv.NewWriter(f).WriteAll(recs)
	if err != nil {
		log.Fatal(err)
	}
}
