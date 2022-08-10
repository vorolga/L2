package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		anagrams map[string][]string
	}{
		{
			"empty arr",
			[]string{},
			map[string][]string{},
		},
		{
			"1 word",
			[]string{"слово"},
			map[string][]string{"слово": {"слово"}},
		},
		{
			"many words",
			[]string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик", "рим", "амамам", "мир", "мама"},
			map[string][]string{"амамам": {"амамам"}, "листок": {"листок", "слиток", "столик"}, "мама": {"мама"},
				"рим": {"рим", "мир"}, "тяпка": {"тяпка", "пятак", "пятка"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test

			result := anagrams(test.words)

			assert.Equal(t, th.anagrams, result)
		})
	}
}
