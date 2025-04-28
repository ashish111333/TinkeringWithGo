package main

import (
	"fmt"
	"time"

	"github.com/ashish111333/twgo/fun"
)

func main() {
	st := time.Now()
	fun.StrWithoutBuffer(100000)
	et := time.Since(st)
	fmt.Println(et)

}
