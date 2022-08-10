package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name      string
		inStr     []string
		inOptions Options
		outStr    []string
		err       error
	}{
		{
			"base sort",
			[]string{"computer",
				"data",
				"debian",
				"laptop",
				"LAPTOP",
				"mouse",
				"RedHat",
				"data",
				"mouse "},
			Options{
				k: -1,
				n: false,
				r: false,
				u: false,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"LAPTOP",
				"RedHat",
				"computer",
				"data",
				"data",
				"debian",
				"laptop",
				"mouse",
				"mouse ",
			},
			nil,
		},
		{
			"base sort",
			[]string{"computer",
				"data",
				"debian",
				"laptop",
				"LAPTOP",
				"mouse",
				"RedHat",
				"data",
				"mouse",
				"mouse    ",
				"mouse  ",
			},
			Options{
				k: -1,
				n: false,
				r: false,
				u: true,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"LAPTOP",
				"RedHat",
				"computer",
				"data",
				"debian",
				"laptop",
				"mouse",
				"mouse  ",
				"mouse    ",
			},
			nil,
		},
		{
			"base sort",
			[]string{"computer",
				"data",
				"debian",
				"laptop",
				"LAPTOP",
				"mouse",
				"RedHat",
				"data",
				"mouse",
				"mouse   ",
				"mouse ",
			},
			Options{
				k: -1,
				n: false,
				r: false,
				u: true,
				M: false,
				b: true,
				c: false,
				h: false,
			},
			[]string{
				"LAPTOP",
				"RedHat",
				"computer",
				"data",
				"debian",
				"laptop",
				"mouse",
				"mouse   ",
				"mouse ",
			},
			nil,
		},
		{
			"base sort",
			[]string{"computer",
				"data",
				"debian",
				"laptop",
				"LAPTOP",
				"mouse",
				"RedHat",
				"data",
				"mouse "},
			Options{
				k: -1,
				n: false,
				r: true,
				u: false,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"mouse ",
				"mouse",
				"laptop",
				"debian",
				"data",
				"data",
				"computer",
				"RedHat",
				"LAPTOP",
			},
			nil,
		},
		{
			"number sort",
			[]string{
				"2",
				"3",
				"1",
				"6",
				"5",
				"4",
				"5",
			},
			Options{
				k: -1,
				n: true,
				r: true,
				u: false,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"6",
				"5",
				"5",
				"4",
				"3",
				"2",
				"1",
			},
			nil,
		},
		{
			"number column sort",
			[]string{
				"2 colorado 322 1M",
				"3 denise 445 2G",
				"1 endrica 100 3m",
				"6 joyce 23 4da",
				"5 marta 40 1T",
				"4 melanie 203 3G",
				"5 marta 40 1T",
			},
			Options{
				k: 2,
				n: true,
				r: false,
				u: true,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"6 joyce 23 4da",
				"5 marta 40 1T",
				"1 endrica 100 3m",
				"4 melanie 203 3G",
				"2 colorado 322 1M",
				"3 denise 445 2G",
			},
			nil,
		},
		{
			"number err sort",
			[]string{
				"2 colorado 322 1M",
				"3 denise 445 2G",
				"1 endrica 100 3m",
				"6 joyce 23 4da",
				"5 marta 40 1T",
				"4 melanie 203 3G",
				"5 marta 40 1T",
			},
			Options{
				k: -1,
				n: true,
				r: false,
				u: false,
				M: false,
				b: false,
				c: false,
				h: false,
			},
			[]string{},
			errors.New("error"),
		},
		{
			"month sort",
			[]string{
				"JULY",
				"AUG",
				"FEB",
				"JAN",
				"march",
			},
			Options{
				k: -1,
				n: false,
				r: false,
				u: false,
				M: true,
				b: false,
				c: false,
				h: false,
			},
			[]string{
				"JAN",
				"FEB",
				"march",
				"JULY",
				"AUG",
			},
			nil,
		},
		{
			"number suf column sort",
			[]string{
				"2 colorado 322 1M",
				"3 denise 445 2G",
				"1 endrica 100 3m",
				"6 joyce 23 4da",
				"5 marta 40 1T",
				"4 melanie 203 3G",
				"5 marta 40 1T",
			},
			Options{
				k: 3,
				n: false,
				r: false,
				u: false,
				M: false,
				b: false,
				c: false,
				h: true,
			},
			[]string{
				"1 endrica 100 3m",
				"6 joyce 23 4da",
				"2 colorado 322 1M",
				"3 denise 445 2G",
				"4 melanie 203 3G",
				"5 marta 40 1T",
				"5 marta 40 1T",
			},
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test

			result, err := sort(th.inStr, th.inOptions)

			if th.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.outStr, result)
			}
		})
	}
}
