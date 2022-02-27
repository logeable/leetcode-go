package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	snums := make([]string, len(nums)-1)
	for i, n := range nums[1:] {
		snums[i] = strconv.Itoa(n)
	}
	return fmt.Sprintf("%d/(%s)", nums[0], strings.Join(snums, "/"))
}
