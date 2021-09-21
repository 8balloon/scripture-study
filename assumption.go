package main

import "strings"

func AssumingDelimiter(delimiter string, text string) []string {
	return strings.Split(text, delimiter)
}

func AssumingPunctuation(punctuation string, text string) []string {
	split := strings.Split(text, punctuation)
	withPunctuation := []string{}
	for _, str := range split {
		withPunctuation = append(withPunctuation, str)
		withPunctuation = append(withPunctuation, punctuation)
	}
	return withPunctuation[:len(withPunctuation)-1]
}
