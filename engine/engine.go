package engine

import (
	"unicode"
)

var vowels = [5][2]rune{
	{'a', '1'},
	{'e', '2'},
	{'i', '3'},
	{'o', '4'},
	{'u', '5'},
}
var vowels2numbers = make(map[rune]rune)
var numbers2vowels = make(map[rune]rune)

func init() {
	for _, pair := range vowels {
		vowels2numbers[pair[0]] = pair[1]
		numbers2vowels[pair[1]] = pair[0]
	}
}

// Encode encodes string according to vowels table
func Encode(input string) (result string) {
	for _, r := range input {
		if val, ok := vowels2numbers[unicode.ToLower(r)]; ok {
			result += string(val)
		} else {
			result += string(r)
		}
	}
	return
}

// Decode decodes encoded string according to vowels table
func Decode(input string) (result string) {
	for _, r := range input {
		if val, ok := numbers2vowels[r]; ok {
			result += string(val)
		} else {
			result += string(r)
		}
	}
	return
}
