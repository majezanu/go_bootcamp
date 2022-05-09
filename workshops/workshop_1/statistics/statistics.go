package statistics

import "fmt"

// Define an ordered list of float numbers and obtain the mean, median and the maximum
// number.
// E.g.
// Input:  [-1, 3, 4, 4, 5]
// Output:  mean = 3
// Median = 4
// maxValue = 5
// —-----------------------------
// Input:  [2, 2, 8, 10]
// Output:  mean = 5.5
// Median = 5
// maxValue = 10

func calculateMean(numbers []int) float64 {
	sum := 0
	size := len(numbers)
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(size)
}

func calculateMaxValue(numbers []int) int {
	size := len(numbers)
	return numbers[size-1]
}

func calculateMedian(numbers []int) float64 {
	size := len(numbers)
	middleIndex := float64(size) / float64(2)
	if size%2 == 0 {
		middleValue1 := numbers[int(middleIndex)]
		middleValue2 := numbers[int(middleIndex)-1]
		numbers := []int{middleValue1, middleValue2}
		return calculateMean(numbers)
	} else {
		median := numbers[int(middleIndex-0.5)]
		return float64(median)
	}
}

func processNumbers(numbers []int) {
	mean := calculateMean(numbers)
	median := calculateMedian(numbers)
	maxValue := calculateMaxValue(numbers)
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Mean: ", mean)
	fmt.Println("Median: ", median)
	fmt.Println("Max Value: ", maxValue)
}

func ProccesNumbers() {
	fmt.Println("Challenge 2: Mean, Median and Max")
	inputs := [][]int{{-1, 3, 4, 4, 5}, {2, 2, 8, 10}}
	for _, input := range inputs {
		processNumbers(input)
	}
}
