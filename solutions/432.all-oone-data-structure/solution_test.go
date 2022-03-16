package solution

import "testing"

func TestAllOne(t *testing.T) {
	all := Constructor()
	all.Inc("a")
	all.Inc("b")
	all.Inc("b")
	all.Inc("c")
	all.Inc("c")
	all.Inc("c")
	all.Dec("b")
	all.Dec("b")
	t.Log(all.GetMinKey())
	all.Dec("a")
	t.Log(all.GetMaxKey())
	t.Log(all.GetMinKey())
}
