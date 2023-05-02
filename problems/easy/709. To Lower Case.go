package easy

import "fmt"

var Alpha = map[string]string{
	"A": "a",
	"B": "b",
	"C": "c",
	"D": "d",
	"E": "e",
	"F": "f",
	"G": "g",
	"H": "h",
	"I": "i",
	"J": "j",
	"K": "k",
	"L": "l",
	"M": "m",
	"N": "n",
	"O": "o",
	"P": "p",
	"Q": "q",
	"R": "r",
	"S": "s",
	"T": "t",
	"U": "u",
	"V": "v",
	"W": "w",
	"X": "x",
	"Y": "y",
	"Z": "z",
}

func ToLowerCase(s string) string {
	str := ""
	for _, v := range s {
		if sa, ok := Alpha[string(v)]; ok {
			str += sa
		} else {
			str += string(v)
		}
	}

	fmt.Println(str)

	return ""
}
