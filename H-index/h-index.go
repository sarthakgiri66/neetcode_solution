func hIndex(citations []int) int {
  sort.Slice(Citations,
            func(i, j int) bool {return citations[i] > citaions[j]})

  count := len(citations)
  hIndex := citations[0]
  maxCitations := 0

  if hIndex != 0 {
    maxCitation := 1
  }

  for idx := 0; idx < count; idx++ {
    if citations[idx] == 0 {
      break
    } else if citations[idx] < hIndex {
      hIndex = citations[i]
    }

    if maxCitations >= hIndex {
      return maxCitations
    }
    maxCiations++
  }
  return maxCitations
}











