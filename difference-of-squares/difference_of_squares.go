package diffsquares

type sum func(int)

func iterate(n int, op sum) {
	for i := 1; i <= n; i++ {
		op(i)
	}
}

//SquareOfSums return the square of the sum of the first n natural numbers.
func SquareOfSums(n int) int {
	s := 0
	iterate(n, func(x int) {
		s += x
	})
	return s * s
}

//SumOfSquares return the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	s := 0
	iterate(n, func(x int) {
		s += x * x
	})
	return s
}

//Difference returns SquareOfSums subtracted from SumOfSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
