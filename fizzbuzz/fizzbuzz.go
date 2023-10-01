package fizzbuzz

import "strconv"

func Run(intInput []int) []string {
	strOutput := make([]string, 0)
	for _, input := range intInput {
		strOutput = append(strOutput, strconv.Itoa(input))
	}
	return strOutput
}
