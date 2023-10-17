package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var count int

func main() {
	var StackB Stack
	var bit_wise Stack
	if len(os.Args) == 2 {
		arg := os.Args[1]
		numbers := strings.Fields(arg)
		StackA := make(Stack, len(numbers))
		for i, numStr := range numbers {
			num, err := parseInt(numStr)
			if err != nil {
				fmt.Println("Error")
				return
			}
			StackA[i] = num
		}
		if CheckDup(StackA) {
			fmt.Println("ERORR")
			return
		} else if sort.IntsAreSorted(StackA) {
			return
		} else {

			if len(numbers) <= 3 {
				SmallStack(StackA)
			} else if len(numbers) >= 4 && len(numbers) <= 6 {
				LargeStack(StackA, StackB)
			} else {

				idx_arr := []int{}
				sortedIndices := countingSort(StackA)
				for _, index := range StackA {
					for idx, num := range sortedIndices {
						if index == num {
							idx_arr = append(idx_arr, idx)
						}
					}
				}

				for _, num := range idx_arr {
					num2 := int64(num)
					num3 := strconv.FormatInt(num2, 2)
					num4, _ := strconv.Atoi(num3)
					bit_wise = append(bit_wise, num4)
				}

				size := len(bit_wise)
				max_num := size - 1
				max_bits := 0
				sort_bits(max_num, max_bits, size, bit_wise, StackB)
				fmt.Println(count - 1)
			}
		}
	} else {
		fmt.Println("Missing input!")
	}
}

func pb(stackA, stackB *[]int) {
	if len(*stackA) == 0 {
		return
	}
	top := (*stackA)[0]
	*stackA = (*stackA)[1:]
	newArray := make([]int, len(*stackB)+1)
	newArray[0] = top
	copy(newArray[1:], *stackB)
	*stackB = newArray
	fmt.Printf("pb\n")
	count++
}

func ra(stackA *[]int) {
	if len(*stackA) == 0 {
		return
	}
	top := (*stackA)[0]
	*stackA = (*stackA)[1:]
	*stackA = append(*stackA, top)
	// Replace with your ra() logic
	fmt.Printf("ra\n")
	count++
}

func pa(stackA, stackB *[]int) {
	if len(*stackB) == 0 {
		return
	}
	top := (*stackB)[0]
	*stackB = (*stackB)[1:]
	newArray := make([]int, len(*stackA)+1)
	newArray[0] = top
	copy(newArray[1:], *stackA)
	*stackA = newArray
	// Replace with your pa() logic
	fmt.Printf("pa\n")
}

func sort_bits(max_num, max_bits, size int, bit_wise, StackB []int) []int {
	for (max_num >> max_bits) != 0 {
		max_bits++
	}
	for i := 0; i < max_bits; i++ {
		for j := 0; j < size; j++ {
			num := bit_wise[0]
			if (num>>i)&1 == 1 {
				ra(&bit_wise)
			} else {
				pb(&bit_wise, &StackB)
			}
		}
		for len(StackB) > 0 {
			pa(&bit_wise, &StackB)
		}
	}
	return bit_wise
}

func countingSort(numbers []int) []int {
	min := findMin(numbers)
	max := findMax(numbers)

	// Create a counting array to store the count of each number
	count := make([]int, max-min+1)

	// Initialize the count array
	for _, num := range numbers {
		count[num-min]++
	}

	// Calculate the prefix sum in the count array
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Create an output array to store the sorted indices
	sortedIndices := make([]int, len(numbers))

	// Build the sorted array using the count array
	for _, num := range numbers {
		index := count[num-min]
		sortedIndices[index-1] = num
		count[num-min]--
	}

	return sortedIndices
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}
