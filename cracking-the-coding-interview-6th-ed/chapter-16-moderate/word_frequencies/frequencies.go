package word_frequencies

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	ErrCannotOpenFile = errors.New("cannot open file")
)

type counter struct {
	words map[string]int
}

func (c counter) Words() map[string]int { return c.words }

func NewCounterFromFile(fileName string) (counter, error) {
	words, err := readWordsFromFile(fileName)
	if err != nil {
		return counter{}, err
	}

	return counter{words}, nil
}

func (c *counter) LoadWordsFromFile(fileName string) error {
	words, err := readWordsFromFile(fileName)
	if err != nil {
		return err
	}
	c.words = words
	return nil
}

func readWordsFromFile(fileName string) (map[string]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, ErrCannotOpenFile
	}
	defer f.Close()

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return nil, err
	}

	words := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		word = string(reg.ReplaceAllString(word, ""))
		word = strings.ToLower(word)

		if v, ok := words[word]; !ok {
			words[word] = 1
		} else {
			words[word] = v + 1
		}
	}

	return words, nil
}

func (c counter) Count(word string) int { return c.words[word] }
