package sevenish

import (
	"fmt"
	"math"
	"slices"
)

// sumOfUniquePowerOf7 returns the sum of the unique power of 7
// We do this with the following criteria
// 1. The number retured must be the smallest unique sum possible
// 2. The number must not be accounted for in the nthNumbers array
func sumOfUniquePowerOf7(sumOfUniquePower int, startMultipleIndex int, lastMultipleIndex int, nthNumbers []int) *int {
	// If the startMultipleIndex is more than the lastMultipleIndex
	// and the sumOfUniquePower seen so far is not in the nthNumbers array,
	// and the lastMultipleIndex is not -1 (it is used as the default last multple index)
	// then we return the sumOfUniquePower seen so far
	if startMultipleIndex > lastMultipleIndex && slices.Index(nthNumbers, sumOfUniquePower) == -1 && lastMultipleIndex != -1 {
		return &sumOfUniquePower
	}

	var result *int = nil

	for i := startMultipleIndex; i <= lastMultipleIndex; i++ {
		multiple := int(math.Pow(7, float64(i)))
		newSumOfUnique := sumOfUniquePower + multiple

		newResult := sumOfUniquePowerOf7(newSumOfUnique, i+1, lastMultipleIndex, nthNumbers)

		// If the result is nil (i.e we have not stored anything in result)
		// However, we computed the newResult (i.e calculated a sum of unique 7s' not present in the nthNumbers array)
		// then we just assign result the value of newResult
		if newResult != nil && result == nil {
			result = newResult

			continue
		}

		// Here we check to make sure that if
		// - newResult != nil and result != nil
		// - and the newResult computed is smaller than result then we assign
		if newResult != nil && result != nil && *newResult < *result {
			result = newResult
		}
	}

	return result
}

func getNthSevenishNumber(lastMultipleIndex int, nthNumbers []int) (result int, isResultAMultiple bool) {
	newSumOfUniquePowerOf7 := sumOfUniquePowerOf7(0, 0, lastMultipleIndex, nthNumbers)
	newMultiple := int(math.Pow(7, float64(lastMultipleIndex+1)))

	// If the sumOfUniquePowerOf7 returns nil (All sums have thus been accounted for) or
	// it's result is greater than just using the multiple, then we use the multiple
	// This is because we want to use a next smallest number possible
	if newSumOfUniquePowerOf7 == nil || *newSumOfUniquePowerOf7 >= newMultiple {
		return newMultiple, true
	}

	return *newSumOfUniquePowerOf7, false
}

func calculateNthSevenishNumber(nthNumber int) int {
	var nthNumbers []int
	var lastMultipleIndex int = -1

	for i := 0; i < nthNumber; i++ {
		result, isResultAMultiple := getNthSevenishNumber(lastMultipleIndex, nthNumbers)
		nthNumbers = append(nthNumbers, result)

		if isResultAMultiple {
			lastMultipleIndex++
		}
	}

	fmt.Printf("The nth numbers are %v\n", nthNumbers)
	return nthNumbers[nthNumber-1]
}

func Init() {
	var nthNumber int

	fmt.Println("=========== Program begins ==============")
	fmt.Println("Enter the nth sevenish number you want to find")

	fmt.Scan(&nthNumber)

	fmt.Printf("The nth sevenish number is %v \n", calculateNthSevenishNumber(nthNumber))
}
