package util

import (
	"fmt"
	"strings"
)

// MaxTagsAllowed default value of allowed tags
var MaxTagsAllowed = 20

// TagsParseAndValidate parses and validates the length of tags
func TagsParseAndValidate(tags string) ([]string, error) {
	tgs := strings.Split(tags, ",")

	if len(tgs) > MaxTagsAllowed {
		return []string{},
			fmt.Errorf("Amount of requested tags is more then allowed")
	}

	if len(tgs) == 1 && tgs[0] == "" {
		return []string{}, nil
	}

	return tgs, nil
}
