package main

import "fmt"

func main() {
	nums := []int{78, 100, 82, 562}
	largest(nums)
}

func finshed() {
	fmt.Println("finshed finding largest")
}

func largest(nums []int) {
	defer finshed()
	fmt.Println("Strated finding largest")
	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("The largest no in", nums, "is", max)
}
