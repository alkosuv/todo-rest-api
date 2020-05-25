package validation

import (
	"regexp"
)

// IsString ...
func IsString(str string, lower int, upper int, pattern string) (bool, error) {
	if len(str) < lower || len(str) > upper {
		return false, ErrLen
	}

	matched, err := regexp.Match(pattern, []byte(str))
	if err != nil {
		return false, err
	}

	if !matched {
		return false, ErrPattern
	}

	return true, nil
}
