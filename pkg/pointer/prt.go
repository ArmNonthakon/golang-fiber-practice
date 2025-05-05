package pointer

func Ptr[T comparable](input T) *T {
	return &input
}
