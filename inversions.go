package inversions

import (
	"io/ioutil"
	"strings"
)

func Inversions(lines []string) (int, error) {
	if len(lines) < 2 {
		return 0, nil
	}
	return 1, nil
}

func FileToInversions(filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")

	return Inversions(lines)
}
