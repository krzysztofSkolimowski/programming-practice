package log

import (
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/ip"
	"github.com/stretchr/testify/assert"
	"net"
	"regexp"
	"testing"
	"time"
)

func Test_checkFilters(t *testing.T) {
	tests := []struct {
		name     string
		log      Log
		filters  []FilterFcn
		expected bool
	}{
		{
			name: "always_true_filter",
			log:  Log{},
			filters: []FilterFcn{
				func(log Log) bool { return true },
			},
			expected: true,
		},
		{
			name: "always_false_filter",
			log:  Log{},
			filters: []FilterFcn{
				func(log Log) bool { return false },
			},
			expected: false,
		},
		{
			name: "chained_always_false_filters",
			log:  Log{},
			filters: []FilterFcn{
				func(log Log) bool { return false },
				func(log Log) bool { return false },
			},
			expected: false,
		},
		{
			name: "chained_always_true_filters",
			log:  Log{},
			filters: []FilterFcn{
				func(log Log) bool { return true },
				func(log Log) bool { return true },
			},
			expected: true,
		},
		{
			name: "chained_always_filters",
			log:  Log{},
			filters: []FilterFcn{
				func(log Log) bool { return true },
				func(log Log) bool { return false },
				func(log Log) bool { return false },
				func(log Log) bool { return true },
			},
			expected: false,
		},
		{
			name: "service_ssh_daemon_true",
			log: Log{
				Service: "sshd",
			},
			filters: []FilterFcn{
				ServiceNameFilterBuilder("sshd"),
			},
			expected: true,
		},
		{
			name: "service_ssh_daemon_false",
			log: Log{
				Service: "CRON",
			},
			filters: []FilterFcn{
				ServiceNameFilterBuilder("sshd"),
			},
			expected: false,
		},
		{
			name: "failed_login_attempt",
			log: Log{
				Msg: "Invalid user admin from 187.12.249.74",
			},
			filters: []FilterFcn{
				FailedLoginAttemptFilter,
			},
			expected: true,
		},
		{
			name: "in_time_range",
			log: Log{
				Time: time.Date(
					0, 11, 30,
					8, 42, 4, 1,
					time.UTC),
			},
			filters: []FilterFcn{
				TimeRangeFilterBuilder(
					time.Date(
						0, 11, 30,
						8, 42, 4, 0,
						time.UTC),
					time.Date(
						0, 11, 31,
						8, 42, 4, 0,
						time.UTC)),
			},
			expected: true,
		},
		{
			name: "in_time_range_false",
			log: Log{
				Time: time.Date(
					0, 11, 29,
					8, 42, 4, 0,
					time.UTC),
			},
			filters: []FilterFcn{
				TimeRangeFilterBuilder(
					time.Date(
						0, 11, 30,
						8, 42, 4, 0,
						time.UTC),
					time.Date(
						0, 11, 31,
						8, 42, 4, 0,
						time.UTC)),
			},
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := checkFilters(tc.log, tc.filters...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIPAddAndFilter(t *testing.T) {
	tests := []struct {
		name             string
		log              Log
		filters          []FilterFcn
		expectedFiltered bool
		expected         []net.IP
	}{
		{
			name: "msg_contains_no_ip_addresses",
			log: Log{
				Msg: "",
			},
			expected: nil,
		},
		{
			name: "msg_contains_ip_addr",
			log: Log{
				Msg: "Nov 30 08:42:04 ip-172-31-27-153 sshd[22182]: Invalid user admin from 187.12.249.7",
			},
			expectedFiltered: true,
			expected: []net.IP{
				net.ParseIP("187.12.249.7"),
			},
		},
		{
			name: "chained_filters",
			log: Log{
				Service: "sshd",
				Msg:     "Nov 30 08:42:04 ip-172-31-27-153 sshd[22182]: Invalid user admin from 187.12.249.7",
			},
			filters: []FilterFcn{
				ServiceNameFilterBuilder("sshd"),
			},
			expectedFiltered: true,
			expected: []net.IP{
				net.ParseIP("187.12.249.7"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := regexp.MustCompile(ip.IPPattern)
			var ipSlice []net.IP = nil
			filters := append(tc.filters, ContainsIPAddAndFilterBuilder(r, &ipSlice))
			actual := checkFilters(tc.log, filters...)
			assert.Equal(t, tc.expectedFiltered, actual)
			assert.Equal(t, tc.expected, ipSlice)
		})
	}
}
