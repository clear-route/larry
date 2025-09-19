package parse

import "strings"

func LastTag(rawTags []byte) string {
	tags := strings.Split(string(rawTags), "\n")
	return tags[len(tags)-2]
}