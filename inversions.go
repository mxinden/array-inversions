package inversions

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Inversions(lines []int) ([]int, int) {
	n := len(lines)
	if n < 2 {
		return lines, 0
	}

	left, leftInv := Inversions(lines[0:(n / 2)])
	right, rightInv := Inversions(lines[(n / 2):])

	combined, splitInv := countSplitInversions(left, right)

	return combined, leftInv + rightInv + splitInv
}

func countSplitInversions(a, b []int) ([]int, int) {
	lenA := len(a)
	aI := 0
	lenB := len(b)
	bI := 0
	c := []int{}
	splitInv := 0

	for k := 0; k < lenA+lenB; k++ {
		if aI >= lenA {
			c = append(c, b[bI])
			bI++
			continue
		}
		if bI >= lenB {
			c = append(c, a[aI])
			aI++
			continue
		}
		if a[aI] < b[bI] {
			c = append(c, a[aI])
			aI++
		} else {
			c = append(c, b[bI])
			bI++
			splitInv += ((lenA+lenB)/2 - aI)
		}
	}

	return c, splitInv
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

	_, inversions := Inversions(intLines)
	return inversions, nil
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
