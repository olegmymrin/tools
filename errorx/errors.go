package errorx

import (
	"log"
)

func FatalOnErr(err error, msgAndArgs ...interface{}) {
	if err != nil {
		log.Fatal(append(msgAndArgs, err)...)
	}
}
