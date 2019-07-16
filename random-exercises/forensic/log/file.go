package log

import (
	"bufio"
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/ip"
	"net"
	"os"
	"regexp"
	"sort"
)

var timeLayout = "Jan 2 15:04:05"

func FromFileFiltered(filePath string, filters ...FilterFcn) ([]Log, error) {
	logs, err := analyzeFile(filePath, filters...)
	if err != nil {
		return nil, err
	}

	return logsToSliceByRowNumber(logs), nil
}

func IPFromFileFiltered(filePath string, filters ...FilterFcn) ([]net.IP, error) {
	var ret []net.IP
	r := regexp.MustCompile(ip.IPPattern)

	f := ContainsIPAddAndFilterBuilder(r, &ret)
	filters = append(filters, f)

	_, err := analyzeFile(filePath, filters...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func analyzeFile(filePath string, filters ...FilterFcn) (map[int]Log, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	parser := NewAuthLogParser(timeLayout)

	logs, n := make(map[int]Log), 0
	for scanner.Scan() {
		line := scanner.Text()

		l, err := parser.Parse(line)
		if err != nil {
			return nil, err
		}

		if checkFilters(l, filters...) {
			logs[n] = l
		}
		n++
	}

	return logs, nil
}

func logsToSliceByRowNumber(logs map[int]Log) []Log {
	keys := sortedKeys(logs)
	ret := make([]Log, len(keys))
	for i, v := range keys {
		ret[i] = logs[v]
	}

	return ret
}

func sortedKeys(logs map[int]Log) []int {
	keys, i := make([]int, len(logs)), 0
	for k := range logs {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	return keys
}
