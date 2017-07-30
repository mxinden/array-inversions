package inversions

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Inversions(lines []int) (int, error) {
	if len(lines) < 2 {
		return 0, nil
	}
	return 1, nil
}

func FileToInversions(filePath string) (int, error) {
	stringLines, err := fileToStringSlice(filePath)
	if err != nil {
		return 0, err
	}

	intLines, err := StringSliceToIntSlice(stringLines)
	if err != nil {
		return 0, err
	}

	return Inversions(intLines)
}

func fileToStringSlice(filePath string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(content), "\n")

	return lines, nil
}

func StringSliceToIntSlice(stringLines []string) ([]int, error) {
	intSlice := []int{}
	for _, s := range stringLines {
		// strings.Split adds an empty string at the end of the slice.
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return intSlice, err
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}
