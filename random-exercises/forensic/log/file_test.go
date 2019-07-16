package log_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFromFileFiltered(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		filters     []log.FilterFcn
		expected    []log.Log
		expectedErr error
	}{
		{
			name: "empty_file_no_filters",
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/empty.log",
			filters:     nil,
			expected:    []log.Log{},
			expectedErr: nil,
		},
		{
			name: "file_is_loaded_no_filters",
			//todo - correct to relative filepath, migrate to modules
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/single.log",
			expected: []log.Log{
				{
					Time: time.Date(
						0, 11, 30,
						6, 39, 00, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "CRON",
					Pid:     21882,
					Msg:     "pam_unix(cron:session): session closed for user root",
				},
			},
			expectedErr: nil,
		},
		{
			name: "file_is_loaded_several_logs_are_parsed_no_filters",
			//todo - correct filepath, migrate to modules
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/short.log",
			expected: []log.Log{
				{
					Time: time.Date(
						0, 11, 30,
						6, 39, 00, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "CRON",
					Pid:     21882,
					Msg:     "pam_unix(cron:session): session closed for user root",
				},
				{
					Time: time.Date(
						0, 11, 30,
						8, 17, 01, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "CRON",
					Pid:     22172,
					Msg:     "pam_unix(cron:session): session closed for user root",
				},
				{
					Time: time.Date(
						0, 11, 30,
						8, 42, 04, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "sshd",
					Pid:     22182,
					Msg:     "Invalid user admin from 187.12.249.74",
				},
			},
			expectedErr: nil,
		},
		{
			name: "service_filter_ssh_daemon",
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/short.log",
			filters: []log.FilterFcn{
				log.ServiceNameFilterBuilder("sshd"),
			},
			expected: []log.Log{
				{
					Time: time.Date(
						0, 11, 30,
						8, 42, 04, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "sshd",
					Pid:     22182,
					Msg:     "Invalid user admin from 187.12.249.74",
				},
			},
			expectedErr: nil,
		},
		{
			name: "service_filter_ssh_daemon_no_results",
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/single.log",
			filters: []log.FilterFcn{
				log.ServiceNameFilterBuilder("sshd"),
			},
			expected:    []log.Log{},
			expectedErr: nil,
		},
		{
			name: "filter_ssh_daemon_and_failed_login_attempt_single_log",
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/attempt.log",
			filters: []log.FilterFcn{
				log.ServiceNameFilterBuilder("sshd"),
				log.FailedLoginAttemptFilter,
			},
			expected: []log.Log{
				{
					Time: time.Date(
						0, 12, 31,
						20, 40, 20, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "sshd",
					Pid:     7857,
					Msg:     "Invalid user oracle from 87.106.50.214",
				},
			},
			expectedErr: nil,
		},
		{
			name: "filter_ssh_daemon_and_failed_login_attempt",
			filePath: "/home/krzysiu/go/src" +
				"/github.com/krzysztofSkolimowski/programming-practice" +
				"/random-exercises/forensic/test-files/test.log",
			filters: []log.FilterFcn{
				log.ServiceNameFilterBuilder("sshd"),
				log.FailedLoginAttemptFilter,
			},
			expected: []log.Log{
				{
					Time: time.Date(
						0, 11, 30,
						8, 42, 4, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "sshd",
					Pid:     22182,
					Msg:     "Invalid user admin from 187.12.249.74",
				},
				{
					Time: time.Date(
						0, 11, 30,
						10, 58, 27, 0,
						time.UTC),
					Name:    "ip-172-31-27-153",
					Service: "sshd",
					Pid:     22291,
					Msg:     "Invalid user admin from 122.225.109.208",
				},
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := log.FromFileFiltered(tc.filePath, tc.filters...)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
