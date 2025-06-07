package main

func main() {

	// fun.StrUsingBuilder is faster than fun.StrWithoutBuilder
	// the builder one starts outperforming the later significantly
	// after the number of strings become large (num of strings >10 appx)
	// question: why does it perform better ? (Think)
	/*err := concurrency.CreateFiles(2, 3, "ash")
	if err != nil {
		fmt.Println(err.Error())
	}*/

	// benchmark AddSLiceItems vs AddSLiceItemsC to see which one if faster does
	// more go routines  always means faster ? Yes or NO ?.

}
