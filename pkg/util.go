package pkg

func NegativeOf[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](value T) T {
	return -value
}
