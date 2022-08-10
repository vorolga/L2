package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err error
	}{
		{
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"abcd",
			"abcd",
			nil,
		},
		{
			"qwe\\4\\5",
			"qwe45",
			nil,
		},
		{
			"qwe\\45",
			"qwe44444",
			nil,
		},
		{
			"45",
			"",
			errors.New("некорректная строка"),
		},
		{
			"",
			"",
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			th := test

			result, err := unpack(test.in)

			if th.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.out, result)
			}
		})
	}
}
