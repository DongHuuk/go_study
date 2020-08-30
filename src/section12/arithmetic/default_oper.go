package arithmetic

type Numbers struct {
	x int64
	y int64
}

func (n *Numbers) SumNumber() int64 {
	return n.x + n.y
}

func (n *Numbers) MinusNumber() int64 {
	return n.x - n.y
}

func (n *Numbers) MultiNumber() int64 {
	return n.x * n.y
}

func (n *Numbers) DivideNumber() int64 {
	return n.x / n.y
}
