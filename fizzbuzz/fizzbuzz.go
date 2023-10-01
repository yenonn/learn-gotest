package fizzbuzz

import "strconv"

func Run(intInput []int) []string {
	strOutput := make([]string, 0)
	for _, input := range intInput {
		if input%3 == 0 {
			strOutput = append(strOutput, "fizz")
		} else {
			strOutput = append(strOutput, strconv.Itoa(input))
		}
	}
	return strOutput
}
