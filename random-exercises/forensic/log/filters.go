package log

import (
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/ip"
	"net"
	"regexp"
	"strings"
	"time"
)

type FilterFcn func(Log) bool

func checkFilters(l Log, filters ...FilterFcn) bool {
	add := true
	for _, f := range filters {
		if !f(l) {
			add = false
			break
		}
	}
	return add
}

func FailedLoginAttemptFilter(l Log) bool { return strings.Contains(l.Msg, "Invalid user") }

func TimeRangeFilterBuilder(start, end time.Time) FilterFcn {
	return func(l Log) bool { return l.Time.After(start) && l.Time.Before(end) }
}

func ServiceNameFilterBuilder(serviceName string) FilterFcn {
	return func(l Log) bool { return l.Service == serviceName }
}

func ContainsIPAddAndFilterBuilder(ipRegexp *regexp.Regexp, ipSlice *[]net.IP) FilterFcn {
	return func(l Log) bool {
		if ipAddr := ip.Find(l.Msg, ipRegexp); ipAddr != nil {
			*ipSlice = append(*ipSlice, ipAddr)
			return true
		}
		return false
	}
}
