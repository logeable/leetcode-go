package solution

import "strconv"

func calPoints(ops []string) int {
	scores := []int{}
	for _, c := range ops {
		switch c {
		case "+":
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		case "D":
			scores = append(scores, scores[len(scores)-1]*2)
		case "C":
			scores = scores[:len(scores)-1]
		default:
			v, _ := strconv.Atoi(c)
			scores = append(scores, v)
		}
	}
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}
