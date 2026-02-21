#Time Complexity: O(n × m)
#n = number of strings
#m = length of the shortest string (the prefix can never be longer than the shortest string
#Space: O(1) — no extra allocations
package main

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, curStr := range strs[1:] {
		if len(prefix) > len(curStr) {
			prefix = prefix[:len(curStr)]
		}
		if len(prefix) > 0 && !strings.HasPrefix(curStr, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}


###Vertical Scanning ✅ (Recommended)###
#Time: O(n × m) — same complexity, but no redundant re-scanning (no HasPrefix re-walk)
#Space: O(1) — no extra allocations
package main

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for i := 0; i < len(strs[0]); i++ {       // iterate columns
        c := strs[0][i]
        for j := 1; j < len(strs); j++ {       // check each string
            if i >= len(strs[j]) || strs[j][i] != c {
                return strs[0][:i]              // mismatch → return what we have
            }
        }
    }
    return strs[0]  // strs[0] itself is the LCP
}


###Sort-based (Most Elegant)###
#Time: O(n log n) for sort + O(m) for comparison
#Space: O(1) extra
#Sorting modifies the original slice. 
package main

import "sort"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    sort.Strings(strs)
    first, last := strs[0], strs[len(strs)-1]

    i := 0
    for i < len(first) && i < len(last) && first[i] == last[i] {
        i++
    }
    return first[:i]
}
