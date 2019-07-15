package log_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		parser        log.Parser
		name          string
		line          string
		expected      log.Log
		expectedError error
	}{
		{
			parser: log.NewAuthLogParser("Jan 2 15:04:05"),
			name:   "test_time_format",
			line:   "Dec 31 22:27:48",
			expected: log.Log{
				Time: time.Date(
					0, 12, 31,
					22, 27, 48, 0,
					time.UTC),
			},
			expectedError: log.ErrIncompleteLogLine,
		},
		{
			parser: log.NewAuthLogParser("Jan 2 15:04:05"),
			name:   "test_parsing_actual_log",
			line: "Dec 31 20:40:43 ip-172-31-27-153 sshd[7907]:" +
				" Received disconnect from 87.106.50.214: 11: Bye Bye [preauth]",
			expected: log.Log{
				Time: time.Date(
					0, 12, 31,
					20, 40, 43, 0,
					time.UTC),
				Name:    "ip-172-31-27-153",
				Service: "sshd",
				Pid:     7907,
				Msg:     "Received disconnect from 87.106.50.214: 11: Bye Bye [preauth]",
			},
			expectedError: nil,
		},
		{
			parser: log.NewAuthLogParser("Jan 2 15:04:05"),
			name:   "test_parsing_incomplete_log_missing_service",
			line:   "Dec 31 20:40:43 ip-172-31-27-153 : qwerty qwerty qwerty",
			expected: log.Log{
				Time: time.Date(
					0, 12, 31,
					20, 40, 43, 0,
					time.UTC),
				Name:    "ip-172-31-27-153",
				Service: "",
				Pid:     0,
				Msg:     "",
			},
			expectedError: log.ErrIncompleteLogLine,
		},
		{
			parser: log.NewAuthLogParser("Jan 2 15:04:05"),
			name:   "test_parsing_incomplete_log_missing_PID",
			line:   "Dec 31 20:40:43 ip-172-31-27-153 sshd: qwerty qwerty qwerty",
			expected: log.Log{
				Time: time.Date(
					0, 12, 31,
					20, 40, 43, 0,
					time.UTC),
				Name:    "ip-172-31-27-153",
				Service: "",
				Pid:     0,
				Msg:     "",
			},
			expectedError: log.ErrIncompleteLogLine,
		},
		{
			parser: log.NewAuthLogParser("Jan 2 15:04:05"),
			name:   "test_parsing_incomplete_log_missing_message",
			line:   "Dec 31 20:40:43 ip-172-31-27-153 sshd[7907]",
			expected: log.Log{
				Time: time.Date(
					0, 12, 31,
					20, 40, 43, 0,
					time.UTC),
				Name:    "ip-172-31-27-153",
				Service: "sshd",
				Pid:     7907,
				Msg:     "",
			},
			expectedError: log.ErrIncompleteLogLine,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.parser.Parse(tc.line)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
