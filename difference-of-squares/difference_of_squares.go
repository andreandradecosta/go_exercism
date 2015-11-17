package diffsquares

//SquareOfSums return the square of the sum of the first n natural numbers.
func SquareOfSums(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s * s
}

//SumOfSquares return the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i * i
	}
	return s
}

//Difference returns SquareOfSums subtracted from SumOfSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
