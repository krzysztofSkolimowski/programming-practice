package log

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrIncompleteLogLine = errors.New("Log line is missing data")
)

type Log struct {
	Time    time.Time
	Name    string
	Service string
	Pid     int
	Msg     string
}

type Parser interface {
	Parse(line string) (Log, error)
}

type authLogParser struct {
	timeLayout string
	//currentYear int
}

func NewAuthLogParser(timeLayout string) *authLogParser {
	return &authLogParser{
		timeLayout: timeLayout,
		//currentYear: currentYear,
	}
}

func (a authLogParser) Parse(line string) (Log, error) {
	ret := Log{}

	//parse time
	t, err := time.Parse(a.timeLayout, line[0:15])
	if err != nil {
		return ret, err
	}
	ret.Time = t
	if len(line) < 16 {
		return ret, ErrIncompleteLogLine
	}
	line = strings.TrimSpace(line[16:])

	//parse name
	i := strings.Index(line, " ")
	if i == -1 {
		return ret, ErrIncompleteLogLine
	}
	ret.Name = line[:i]
	if len(line) < 1 {
		return ret, ErrIncompleteLogLine
	}
	line = strings.TrimSpace(line[i:])

	//parse service
	i = strings.Index(line, "[")
	if i == -1 {
		return ret, ErrIncompleteLogLine
	}
	ret.Service = line[:i]
	line = strings.TrimSpace(line[i+1:])

	//parse pid
	i = strings.Index(line, "]")
	if i == -1 {
		return ret, ErrIncompleteLogLine
	}
	ret.Pid, err = strconv.Atoi(line[:i])
	if err != nil {
		return Log{}, errors.Wrap(err, "cannot parse pid")
	}

	//parse msg
	if len(line) < i+2 {
		return ret, ErrIncompleteLogLine
	}
	line = line[i+2:]
	ret.Msg = strings.TrimSpace(line)

	return ret, nil
}

//func (a authLogParser) Parse(line string) (Log, error) {
//	ret := Log{}
//
//	t, err := time.Parse(a.timeLayout, line[0:15])
//	if err != nil {
//		return ret, err
//	}
//
//	t = t.AddDate(a.currentYear, 0, 0)
//	ret.Time = t
//
//	if len(line) < 16 {
//		return ret, ErrIncompleteLogLine
//	}
//
//	line = strings.TrimSpace(line[16:])
//	fields := strings.Fields(line)
//
//	if len(fields) >= 2 {
//		ret.Name = fields[0]
//
//		serviceWithPid := strings.Split(fields[1], "[")
//		if len(serviceWithPid) < 2 {
//			return ret, ErrIncompleteLogLine
//		}
//
//		ret.Service = serviceWithPid[0]
//		pidStr := strings.Trim(serviceWithPid[1], "]:")
//
//		ret.Pid, err = strconv.Atoi(pidStr)
//		if err != nil {
//			return ret, ErrIncompleteLogLine
//		}
//
//		ret.Msg = strings.Join(fields[2:], " ")
//		if len(ret.Msg) == 0 {
//			return ret, ErrIncompleteLogLine
//		}
//	}
//
//	return ret, nil
//}
