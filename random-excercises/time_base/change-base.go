package time_base

import "strconv"

type timeUnit struct {
	name                 string
	base, divisionFactor int
}

var (
	timeUnits = []timeUnit{
		{"sec", 60, 1},
		{"min", 60, 60},
		{"h", 60, 3600},
		{"days", 24, 86400},
	}
)

func ChangeTime(n int) string {
	result := ""

	for _, u := range timeUnits {
		n, result = parseTimeUnit(n, result, u.name, u.base, u.divisionFactor)
	}

	return result
}

//todo - check negative
func parseTimeUnit(n int, parsedTime string, displayName string, base int, divisionFactor int) (int, string) {
	seconds := (n / divisionFactor) % base
	if seconds > 0 {
		n = n - seconds
		parsedTime = " " + strconv.Itoa(seconds) + " " + displayName + parsedTime
	}
	return n, parsedTime
}
