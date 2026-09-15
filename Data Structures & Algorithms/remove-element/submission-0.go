func removeElement(nums []int, val int) int {
	writePosition := 0
	for i := range nums {
		if nums[i] != val {
			nums[writePosition] = nums[i]
			writePosition++
		}
	}
	return writePosition
}
