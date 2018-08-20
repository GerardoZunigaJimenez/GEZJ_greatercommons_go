package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	sParsed := removeSpecialCharacters(s)
	sIgnoringCase := strings.ToLower(sParsed)
	xs := strings.Fields(sIgnoringCase)

	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//CountWordsByString returns the number of words by the string.
// A single space is used to identify the words
func CountWordsByString(s string) int {
	sParsed := removeSpecialCharacters(s)

	var i = 0
	for i, _ = range strings.Split(sParsed, " ") {
		i++
	}
	return i
}

//removeSpecialCharacters will return the original string without
// empty single spaces, double quotes, commas, semicolons and dots.
func removeSpecialCharacters(s string) string {
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ";", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "?", "", -1)
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, ":", "", -1)
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Replace(s, "-", " ", -1)
	return s
}
