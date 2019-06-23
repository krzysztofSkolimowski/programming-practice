package english_int

import (
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrNumberTooLarge = errors.New("Number is too large to convert")
)

// English Int: Given any integer, print an
// English phrase that describes the integer
// (e.g., "One Thousand, Two Hundred Thirty Four)
func Convert(n int) (string, error) {
	return convert(n)
}

var (
	ones = [...]string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}
	teens = [...]string{
		"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}
	tens = [...]string{
		"ten", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety",
	}

	larges = [...] string{
		"", "thousand", "million", "billion",
	}
)

const base = 1000

func convert(n int) (string, error) {
	if n == 0 {
		return ones[0], nil
	}

	ret := ""
	for i := range larges {

		num := n % base
		numString, err := convertSmall(num)

		if err != nil {
			return "", err
		}

		if num != 0 {
			ret = numString + " " + larges[i] + " " + ret
		}

		if n < base {
			break
		}

		n = (n - num) / base
	}

	return strings.TrimSpace(ret), nil
}

//converts number from 0 - 999
func convertSmall(n int) (string, error) {
	switch {
	case n < 10:
		return ones[n], nil

	case n < 19:
		return teens[n-10], nil

	case n < 99:
		ones := n % 10
		if ones == 0 {
			return tens[n/10-1], nil
		} else {
			tensConverted, err := convertSmall(n - ones)
			if err != nil {
				return "", err
			}
			onesConverted, err := convertSmall(ones)
			if err != nil {
				return "", err
			}
			return tensConverted + " " + onesConverted, nil
		}

	case n < 999:
		tens := n % 100
		if tens == 0 {
			return ones[n/100] + " hundred", nil
		} else {
			tensConverted, err := convertSmall(tens)
			if err != nil {
				return "", err
			}
			return ones[n/100] + " hundred " + tensConverted, nil
		}

	default:
		return "", ErrNumberTooLarge
	}
}
