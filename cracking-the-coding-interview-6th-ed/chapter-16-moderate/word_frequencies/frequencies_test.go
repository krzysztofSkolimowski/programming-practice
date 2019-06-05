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
			name:        "file-is-opened",
			fileName:    "lorem",
			expectedErr: nil,
		},
		{
			name:        "file-is-not-opened",
			fileName:    "loremeee",
			expectedErr: word_frequencies.ErrCannotOpenFile,
		},
		{
			name:     "file-is-not-opened",
			fileName: "test-book-1",
			expectedWordsCount: map[string]int{
				"abc": 4,
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

//func Test_counter_Count(t *testing.T) {
//	type fields struct {
//		words map[string]int
//	}
//	type args struct {
//		word string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := counter{
//				words: tt.fields.words,
//			}
//			if got := c.Count(tt.args.word); got != tt.want {
//				t.Errorf("counter.Count() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
