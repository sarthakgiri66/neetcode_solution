func productExceptSelf(nums []int) []int {
  count := len(nums)
  product := make([]int, count)

  prefix := 1
  for i := 0; i < count; i++ {
    product[i] = prefix
    prefix *= nums[i]
  }

  suffix := 1
  for i := count - 1; i >= 0; i-- {
    product[i] *= suffix
    suffix *= nums[i]
  }

  return product
}
