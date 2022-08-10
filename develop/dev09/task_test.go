package main

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	tests := []struct {
		filename string
		url      string
		err      error
	}{
		{
			"index.html",
			"http://www.site.com",
			nil,
		},
		{
			"robots.txt",
			"http://www.google.com/robots.txt",
			nil,
		},
		{
			"index.html",
			"http://asdaa",
			errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			th := test

			err := downloadFile(th.filename, th.url)

			if th.err != nil {
				assert.Error(t, err)
				err = os.Remove(th.filename)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				assert.NoError(t, err)
				_, err = os.Stat(th.filename)
				assert.NoError(t, err)
				err = os.Remove(th.filename)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}
