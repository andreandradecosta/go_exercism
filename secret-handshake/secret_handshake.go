package secret

import "strconv"

var handshakes = [4]string{"wink", "double blink", "close your eyes", "jump"}

const reverse = 4

func Handshake(n int) []string {
	var result []string
	if n >= 0 {
		binReversed := reverseString(strconv.FormatInt(int64(n), 2))
		for i := 0; i < len(binReversed) && i < len(handshakes); i++ {
			if string(binReversed[i]) == "1" {
				result = append(result, handshakes[i])
			}
		}
		if len(binReversed) > reverse && string(binReversed[reverse]) == "1" {
			result = reverseSlice(result)
		}
	}
	return result
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
