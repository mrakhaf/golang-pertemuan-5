package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameChecker(t *testing.T) {

	testcases := []struct {
		param    string
		expected string
	}{
		{
			param:    "Abdul",
			expected: "Found Abdul!",
		},
		{
			param:    "Rakha",
			expected: "Not Abdul!",
		},
	}

	for _, testcase := range testcases {
		result := NameChecker(testcase.param)
		assert.Equal(t, testcase.expected, result)
	}

}
