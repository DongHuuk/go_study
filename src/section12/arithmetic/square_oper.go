package arithmetic

func (n *Numbers) Plus() int64 {
	return n.x*n.x + n.y*n.y
}

func (n *Numbers) Minus() int64 {
	return n.x*n.x - n.y*n.y
}
