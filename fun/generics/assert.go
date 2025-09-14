package generics

func assertEqual[T comparable](a, b T) bool {
	return a == b
}
