package regexlite

type number interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type nWrapper[T number] struct {
	value T
	err   error
}

type sWrapper struct {
	value string
	err   error
}
