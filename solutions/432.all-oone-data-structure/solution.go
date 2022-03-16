package solution

import (
	"container/list"
	"fmt"
)

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

type node struct {
	keys  map[string]struct{}
	count int
}

func (n *node) String() string {
	var keys []string
	for k := range n.keys {
		keys = append(keys, k)
	}
	return fmt.Sprintf("count: %d, keys: %q", n.count, keys)
}

func Constructor() AllOne {
	return AllOne{
		List:  list.New(),
		nodes: make(map[string]*list.Element),
	}
}

func (this *AllOne) Inc(key string) {
	ele, ok := this.nodes[key]
	if !ok {
		if this.Front() == nil || this.Front().Value.(*node).count > 1 {
			ele = this.PushFront(&node{
				count: 1,
				keys: map[string]struct{}{
					key: {},
				},
			})
		} else {
			this.Front().Value.(*node).keys[key] = struct{}{}
			ele = this.Front()
		}
	} else {
		n := ele.Value.(*node)
		delete(n.keys, key)
		oldEle := ele

		nextEle := ele.Next()
		if nextEle == nil || nextEle.Value.(*node).count > n.count+1 {
			ele = this.InsertAfter(&node{
				count: n.count + 1,
				keys: map[string]struct{}{
					key: {},
				},
			}, ele)
		} else {
			nextEle.Value.(*node).keys[key] = struct{}{}
			ele = nextEle
		}
		if len(n.keys) == 0 {
			this.Remove(oldEle)
		}
	}
	this.nodes[key] = ele
}

func (this *AllOne) Dec(key string) {
	ele, ok := this.nodes[key]
	if ok {
		n := ele.Value.(*node)
		delete(n.keys, key)
		oldEle := ele
		if n.count > 1 {
			prevEle := ele.Prev()
			if prevEle == nil || prevEle.Value.(*node).count < n.count-1 {
				ele = this.InsertBefore(&node{
					count: n.count - 1,
					keys: map[string]struct{}{
						key: {},
					},
				}, ele)
			} else {
				prevEle.Value.(*node).keys[key] = struct{}{}
				ele = prevEle
			}
			this.nodes[key] = ele
		} else {
			delete(this.nodes, key)
		}
		if len(n.keys) == 0 {
			this.Remove(oldEle)
		}

	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Back() == nil {
		return ""
	}
	keys := (this.Back().Value.(*node)).keys
	for k := range keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Front() == nil {
		return ""
	}
	keys := (this.Front().Value.(*node)).keys
	for k := range keys {
		return k
	}
	return ""
}
