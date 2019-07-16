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
}

func NewAuthLogParser(timeLayout string) *authLogParser {
	return &authLogParser{
		timeLayout: timeLayout,
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
