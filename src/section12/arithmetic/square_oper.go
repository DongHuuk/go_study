package arithmetic

func (n *Numbers) Plus() int64 {
	return n.X*n.X + n.Y*n.Y
}

func (n *Numbers) Minus() int64 {
	return n.X*n.X - n.Y*n.Y
}
