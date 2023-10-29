package main

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapsort(nums, n, i)
	}
	for i := n - 1; i >= n-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsort(nums, i, 0)
	}
	return nums[0]
}

func heapsort(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapsort(arr, n, largest)
	}
}
