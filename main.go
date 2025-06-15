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

	// benchmark AddSLiceItems vs AddSLiceItemsC vs AddSliceItemsCChannels vs AddSlice see
	// which one is faster make observation about below points.

	// 1. more go routines  always means faster ? Yes or NO ?.

	// 2. fewer atomic operation better than more ? Yes or NO ?.

	// 3. in the above example for small slices the concurrent version
	//    is always slower than the synchronous version (but how small we are talking Think !!),
	//    do you face similar problem when slice is really large  ?.

	// 4. Important observation:- the the addSliceItems using channels is faster
	//	  than the ones using atomics ...why ?????

	//5.  AddSliceItemsCMx is even faster than the channels version of AddSliceItems
	//	  final conclusion-------->
	//	  AddSliceItemsCMx	> AddSliceItemsCChannels > AddSliceItemsC > AddSliceItems
}
