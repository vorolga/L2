package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		name       string
		resultTime time.Time
		error      error
	}{
		{
			name:       "CorrectAnswer",
			resultTime: time.Now(),
			error:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			currentTime, err := getCurrentTime()

			assert.NoError(t, err)
			assert.Condition(t, func() bool {
				return currentTime.Sub(th.resultTime) < time.Second
			})
		})
	}
}
