package generics

// generic function that compares two types
func sum[T int | string | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum

}
