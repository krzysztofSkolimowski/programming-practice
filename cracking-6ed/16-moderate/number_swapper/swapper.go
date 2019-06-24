package number_swapper

/* Number Swapper: Write a function to swap
a number in place (that is, without temporary variables) */
func Swap(a, b *int) {
	*b = *b - *a
	*a = *a + *b
	*b = *a - *b
	return
}
