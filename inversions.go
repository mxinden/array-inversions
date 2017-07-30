package inversions

import (
	"io/ioutil"
	"strings"
)

func Inversions(lines []string) (int, error) {
	return 0, nil
}

func FileToInversions(filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {

	}
	lines := strings.Split(string(content), "/n")
	return Inversions(lines)
}
