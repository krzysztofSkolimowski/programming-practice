package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func AnalyzeFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	//fInfo, err := f.Stat()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("name: %s \t size: %d \t mode: %d \t time: %v \t isDir: %v\n", fInfo.Name(), fInfo.Size(), fInfo.Mode(), fInfo.ModTime(), fInfo.IsDir())

	timeLayout := "Jan 2 15:04:05"
	scanner := bufio.NewScanner(f)
	logs := make(map[int]Log)

	parser := NewAuthLogParser(timeLayout)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		l, err := parser.Parse(line)
		if err != nil {
			log.Println("ERROR: ", err)
		}

		logs[i] = l
		i++
	}

	for j := 0; j < i; j++ {
		v, ok := logs[j]
		if !ok {
			continue
		}

		fmt.Printf("%d. \t %v \t %s \t %s \t %d \t %s \n", j, v.Time, v.Name, v.Service, v.Pid, v.Msg)
	}

	return nil
}
