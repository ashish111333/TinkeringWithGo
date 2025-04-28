package fun

// creates a new string "there" in each iteration
// and appends it to gibberishMssg
func StrWithoutBuffer(iterations int) {

	var gibberishMssg = "hello"
	for i := 0; i < iterations; i++ {
		gibberishMssg += "there"
	}

}

func StrWithBuffer(iterations int) {

}
