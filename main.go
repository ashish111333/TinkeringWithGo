package main

import (
	"fmt"

	"github.com/ashish111333/twgo/concurrency"
)

func main() {

	// fun.StrUsingBuilder is faster than fun.StrWithoutBuilder
	// the builder one starts outperforming the later significantly
	// after the number of strings become large (num of strings >10 appx)
	// question: why does it perform better ? (Think)
	/*err := concurrency.CreateFiles(2, 3, "ash")
	if err != nil {
		fmt.Println(err.Error())
	}*/
	s := concurrency.RandIntSlice(2, 5)
	sum := concurrency.AddSliceItemsC(2, s)
	fmt.Println(sum)
}
