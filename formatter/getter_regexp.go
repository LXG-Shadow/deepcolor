package formatter

import (
	"regexp"
)

type regexpValueGetter struct {
	data string
}

func (r *regexpValueGetter) Get(expression string) interface{} {
	re, err := regexp.Compile(expression)
	if err != nil {
		return nil
	}

	// Find all matches
	matches := re.FindAllStringSubmatch(r.data, -1)
	if len(matches) == 0 {
		// No match found
		return nil
	}

	// If there's more than one match, return an array of matches
	if len(matches) > 1 {
		results := make([]interface{}, 0, len(matches))
		for _, m := range matches {
			if len(m) > 1 {
				// If there's a capturing group, return the first one
				results = append(results, m[1])
			} else {
				// Otherwise return the entire match
				results = append(results, m[0])
			}
		}
		return results
	}

	// Single match found
	if len(matches[0]) > 1 {
		// If there's a capturing group, return the first one
		return matches[0][1]
	}
	// Otherwise return the entire match
	return matches[0][0]
}

func NewRegexpValueGetter(data string) IValueGetter {
	return &regexpValueGetter{
		data: data,
	}
}
