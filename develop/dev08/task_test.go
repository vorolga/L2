package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecCmd(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{
			"cd",
			errors.New("error"),
		},
		{
			"abc",
			errors.New("error"),
		},
		{
			"pwd",
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			th := test

			err := execCmd(th.input)

			if th.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
