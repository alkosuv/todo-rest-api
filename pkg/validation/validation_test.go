package validation_test

import (
	"testing"

	"github.com/gen95mis/todo-rest-api/pkg/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidation_IsString(t *testing.T) {
	testCase := []struct {
		name          string
		str           string
		lower         int
		upper         int
		pattern       string
		isValid       bool
		expectedError error
	}{
		{
			name:          "valid",
			str:           "string",
			lower:         5,
			upper:         15,
			pattern:       `\w`,
			isValid:       true,
			expectedError: nil,
		},
		{
			name:          "err lower",
			str:           "",
			lower:         5,
			upper:         15,
			pattern:       `\w`,
			isValid:       false,
			expectedError: validation.ErrLen,
		},
		{
			name:          "err upper",
			str:           "1234567",
			lower:         5,
			upper:         6,
			pattern:       `\w`,
			isValid:       false,
			expectedError: validation.ErrLen,
		},
		{
			name:          "err pattern",
			str:           "*@#$%^!",
			lower:         5,
			upper:         15,
			pattern:       `\w`,
			isValid:       false,
			expectedError: validation.ErrPattern,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := validation.IsString(tc.str, tc.lower, tc.upper, tc.pattern)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.isValid, valid)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
			}
		})
	}

}
