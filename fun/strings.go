package fun

import (
	"strings"
)

// creates a new string "there" in each iteration
// and appends it to gibberishMssg string is created in heap
func StrWithoutBuilder(iterations int) {

	var gibberishMssg = "hello"
	for i := 0; i < iterations; i++ {
		gibberishMssg += "there"
	}

}

// uses string builder to create string
func StrUsingBuilder(iterations int) {
	sb := strings.Builder{}
	sb.WriteString("hello")
	for i := 0; i < iterations; i++ {
		sb.WriteString("there")
	}

}
