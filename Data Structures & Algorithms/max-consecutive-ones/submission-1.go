func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	currentNumber := 0
	
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentNumber += 1
			max = int(math.Max(float64(max), float64(currentNumber)))
		} else {
			currentNumber = 0
		}
	}
	return max
}
