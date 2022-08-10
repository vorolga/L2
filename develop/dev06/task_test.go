package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name      string
		inStr     []string
		inOptions *Options
		outStr    []string
	}{
		{
			"cut f",
			[]string{"Name\t:Company:Price:Ganre:Multiplayer", "Item1\t:Company1:60000$:Ganre1:Yes",
				"Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4\t:Company4:90000$:Ganre1:Yes", "Item5\t:Company5:110000$:Ganre1:Yes",
				"Item6\t:Company6:410000$:Ganre2:Yes"},
			&Options{
				f: 1,
				d: "\t",
				s: false,
			},
			[]string{"Name", "Item1", "Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4", "Item5", "Item6"},
		},
		{
			"cut f s",
			[]string{"Name\t:Company:Price:Ganre:Multiplayer", "Item1\t:Company1:60000$:Ganre1:Yes",
				"Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4\t:Company4:90000$:Ganre1:Yes", "Item5\t:Company5:110000$:Ganre1:Yes",
				"Item6\t:Company6:410000$:Ganre2:Yes"},
			&Options{
				f: 1,
				d: "\t",
				s: true,
			},
			[]string{"Name", "Item1", "Item4", "Item5", "Item6"},
		},
		{
			"cut big f s",
			[]string{"Name\t:Company:Price:Ganre:Multiplayer", "Item1\t:Company1:60000$:Ganre1:Yes",
				"Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4\t:Company4:90000$:Ganre1:Yes", "Item5\t:Company5:110000$:Ganre1:Yes",
				"Item6\t:Company6:410000$:Ganre2:Yes"},
			&Options{
				f: 5,
				d: "\t",
				s: true,
			},
			[]string{"", "", "", "", ""},
		},
		{
			"cut big f",
			[]string{"Name\t:Company:Price:Ganre:Multiplayer", "Item1\t:Company1:60000$:Ganre1:Yes",
				"Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4\t:Company4:90000$:Ganre1:Yes", "Item5\t:Company5:110000$:Ganre1:Yes",
				"Item6\t:Company6:410000$:Ganre2:Yes"},
			&Options{
				f: 5,
				d: "\t",
				s: false,
			},
			[]string{"", "", "Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes", "", "", ""},
		},
		{
			"cut f d",
			[]string{"Name\t:Company:Price:Ganre:Multiplayer", "Item1\t:Company1:60000$:Ganre1:Yes",
				"Item2:Company2:70000$:Ganre2:Yes", "Item3:Company3:90000$:Ganre1:Yes",
				"Item4\t:Company4:90000$:Ganre1:Yes", "Item5\t:Company5:110000$:Ganre1:Yes",
				"Item6\t:Company6:410000$:Ganre2:Yes"},
			&Options{
				f: 2,
				d: ":",
				s: false,
			},
			[]string{"Company", "Company1", "Company2", "Company3", "Company4", "Company5", "Company6"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test

			result := cut(th.inOptions, th.inStr)

			assert.Equal(t, th.outStr, result)
		})
	}
}
