func canJump(nums []int) bool{
    maxJumpLen := 0 
    retVal := true 

    for _, curJumpLen := range nums {
      if maxJumpLen < 0 {
        retVal = false
        break
      } else if curJumpLen > maxJumpLen {
          maxJumpLen = curJumpLen
      }
      maxJumpLen--
    }
    return retVal
}
