package solution

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	cache := make(map[int]*Employee)
	for _, emp := range employees {
		cache[emp.Id] = emp
	}

	total := 0

	queue := []int{id}

	for len(queue) > 0 {
		cur := cache[queue[0]]
		queue = queue[1:]

		total += cur.Importance
		queue = append(queue, cur.Subordinates...)
	}

	return total
}
