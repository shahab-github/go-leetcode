package main

import "fmt"

func twoSum(nums []int, target int) []int {

    numIndices := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, found := numIndices[complement]; found {
            return []int{index, i}
        }

        numIndices[num] = i
    }

    return []int{}
}

func main() {
    nums1 := []int{2, 7, 11, 15}
    target1 := 18
    fmt.Println(twoSum(nums1, target1))  // Output: [0 1]

    nums2 := []int{3, 2, 4}
    target2 := 6
    fmt.Println(twoSum(nums2, target2))  // Output: [1 2]

    nums3 := []int{3, 3}
    target3 := 6
    fmt.Println(twoSum(nums3, target3))  // Output: [0 1]
}
