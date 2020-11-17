package parking

//priority queue item
type PriorityQueue []uint

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(uint))

}

func (p *PriorityQueue) Pop() interface{} {
	n := len(*p)
	old := (*p)[n-1]
	*p = (*p)[0 : n-1]
	return old
}

func (p PriorityQueue) Len() int           { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i] < p[j] }
func (p PriorityQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
