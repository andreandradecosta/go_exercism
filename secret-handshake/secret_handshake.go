package secret

import "strconv"

var handshakes = [4]string{"wink", "double blink", "close your eyes", "jump"}

const reverse = 5

func Handshake(n int) []string {
	var result []string
	if n > 0 {
		binReversed := strconv.FormatInt(int64(n), 2)
		for i, j := len(binReversed)-1, 0; i >= 0 && j < len(handshakes); i, j = i-1, j+1 {
			if string(binReversed[i]) == "1" {
				result = append(result, handshakes[j])
			}
		}
		if len(binReversed) >= reverse && string(binReversed[len(binReversed)-reverse]) == "1" {
			result = reverseSlice(result)
		}
	}
	return result
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
