package validation

import (
	"regexp"
)

// IsString ...
func IsString(str string, pattern string) (bool, error) {
	matched, err := regexp.Match(pattern, []byte(str))
	if err != nil {
		return false, err
	}

	if !matched {
		return false, ErrPattern
	}

	return true, nil
}
