package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	recs, err := csv.NewReader(f).ReadAll()
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
	f, err = os.OpenFile(os.Args[2], os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(recs)
	if err != nil {
		log.Fatal(err)
	}
}
