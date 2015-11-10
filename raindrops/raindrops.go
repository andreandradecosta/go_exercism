// +build !example

package raindrops

import (
	"fmt"
	"sort"
	"strings"
)

//TestVersion checks if this code is running on the right test.
const TestVersion = 1

var primes = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

var primeKeys []int

func init() {
	for k := range primes {
		primeKeys = append(primeKeys, k)
	}
	sort.Ints(primeKeys)
}

// Convert returns a string representing the number prime-factorization:
// if the number contains 3 as a prime factor, output 'Pling';
// if the number contains 5 as a prime factor, output 'Plang';
// if the number contains 7 as a prime factor, output 'Plong';
// if the number does not contain 3, 5, or 7 as a prime factor,
// just pass the number's digits straight through.
func Convert(number int) string {
	var s []string
	n := number

	for _, p := range primeKeys {
		if n%p == 0 {
			s = append(s, primes[p])
			n /= p
		}
	}
	if len(s) == 0 {
		return fmt.Sprintf("%v", number)
	}
	return strings.Join(s, "")
}

// The test program has a benchmark too.  How fast does your Convert convert?
//BenchmarkConvert          300000              4533 ns/op
