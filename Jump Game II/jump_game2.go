func jump(nums []int) int {
  jumpCount := 0
  lastMaxReach := 0
  curMaxReach := 0
  count := len(nums)
  if count == 1 {
    return jumpCount
  }

  for i := 0; i < count; i++ {
    curMaxReach = max(curMaxReach, i + nums[i])
    if i == lastMaxReach {
      jumpCount++
      lastMaxReach = curMaxReach
      if lastMaxReach >= count - 1 {
        break
      }
    }
  }
  return jumpCount
}
