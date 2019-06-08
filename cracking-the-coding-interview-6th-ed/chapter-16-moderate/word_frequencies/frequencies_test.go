package word_frequencies_test

import (
	"fmt"
	"github.com/krzysztofSkolimowski/programming-practice/cracking-the-coding-interview-6th-ed/chapter-16-moderate/word_frequencies"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCounterFromFile(t *testing.T) {
	tests := []struct {
		name               string
		fileName           string
		expectedWordsCount map[string]int
		expectedErr        error
	}{
		{
			name:        "file_is_opened",
			fileName:    "lorem",
			expectedErr: nil,
		},
		{
			name:        "file_is_not_opened",
			fileName:    "loremeee",
			expectedErr: word_frequencies.ErrCannotOpenFile,
		},
		{
			name:     "count_works_correctly",
			fileName: "test-book-1",
			expectedWordsCount: map[string]int{
				"abc": 4,
			},
			expectedErr: nil,
		},
		{
			name:     "count_works_correctly_many_Words",
			fileName: "test-book-2",
			expectedWordsCount: map[string]int{
				"abc": 4,
				"aa":  3,
				"asd": 2,
			},
			expectedErr: nil,
		},
		{
			name:     "count_works_correctly_multiple_lines",
			fileName: "test-book-3",
			expectedWordsCount: map[string]int{
				"abc": 12,
				"aa":  9,
				"asd": 6,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			counter, err := word_frequencies.NewCounterFromFile(tc.fileName)
			assert.Equal(t, tc.expectedErr, err)
			fmt.Println(counter.Words())
			if tc.expectedWordsCount != nil {
				for k, v := range tc.expectedWordsCount {
					assert.Equal(t, v, counter.Count(k))
				}
			}
		})
	}
}