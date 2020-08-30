package arithmetic

type Numbers struct {
	X int64
	Y int64
}

func (n *Numbers) SumNumber() int64 {
	return n.X + n.Y
}

func (n *Numbers) MinusNumber() int64 {
	return n.X - n.Y
}

func (n *Numbers) MultiNumber() int64 {
	return n.X * n.Y
}

func (n *Numbers) DivideNumber() int64 {
	return n.X / n.Y
}
