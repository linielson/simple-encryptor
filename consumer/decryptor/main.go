package decryptor

import "strings"

func DecryptMessage(msg string) string {
	return decode(msg)
}

func decode(msg string) string {
	replacer := strings.NewReplacer(
		"/4", "a", "/3", "e", "/1", "i", "/0", "o", "/8", "b", "/6", "c", "/9", "g", "/5", "s", "/7", "t", "/2", "z",
	)
	return replacer.Replace(msg)
}
