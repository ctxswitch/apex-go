package apex

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameBuilder(t *testing.T) {
	tests := []struct {
		ns       string
		sub      string
		name     string
		sep      rune
		expected string
		err      error
	}{
		{"apex", "test", "counter", '_', "apex_test_counter", nil},
		{"apex", "", "counter", '_', "apex_counter", nil},
		{"", "test", "counter", '_', "test_counter", nil},
		{"", "", "counter", '_', "counter", nil},
		{"apex", "test", "", '_', "", ErrInvalidMetricName},
	}

	for _, tt := range tests {
		retval, err := NameBuilder(tt.ns, tt.sub, tt.name, tt.sep)
		if tt.err != nil {
			assert.ErrorIs(t, err, tt.err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, retval)
		}
	}
}

func TestLabelKeys(t *testing.T) {
	retval := Labels{
		"one": "1",
		"two": "2",
	}.Keys()
	sort.Strings(retval)
	expected := []string{"one", "two"}
	assert.Equal(t, expected, retval)
}
