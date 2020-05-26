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
		pattern       string
		isValid       bool
		expectedError error
	}{
		{
			name:          "valid",
			str:           "string",
			pattern:       `\w`,
			isValid:       true,
			expectedError: nil,
		},
		{
			name:          "err pattern",
			str:           "*@#$%^!",
			pattern:       `\w`,
			isValid:       false,
			expectedError: validation.ErrPattern,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := validation.IsString(tc.str, tc.pattern)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.isValid, valid)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
			}
		})
	}

}
