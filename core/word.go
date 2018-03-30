package core

import (
	"unicode"
	"strings"
)

func GetNameBySnakeCase(in string)string{
	runes := []rune(in)
	var out []rune
	for i,w := range runes{
		switch  {
		case i == 0:
			out = append(out ,unicode.ToUpper(w))
		case out[i-1] == ' ':
			out = append(out ,unicode.ToUpper(w))
		case w == '_':
			out = append(out,' ')
		default:
			out = append(out,w)
		}
	}
	return string(out)
}

func RemoveExtension(str string)string{
	index := strings.LastIndexAny(str,".")
	return str[:index]
}