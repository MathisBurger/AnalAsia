package algorithms

import "strings"

func RemovePunicationMarks(word string) string {

	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(word, ",", ""),
					"!", ""),
				".", ""),
			"?", ""),
		"\t", "")
}
