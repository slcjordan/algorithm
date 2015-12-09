package main

type color bool

const (
	black color = iota == 0
	red
)

type initialized struct {
	left  RBNode
	right RBNode
	value string
}

type RBNode struct {
	color color
	*initialized
}

func (n RBNode) Insert(v string) RBNode {
	if n.initialized == nil {
		n.initialized = &initialized{value: v}
		n.color = red
		return n
	}
	if n.value == v {
		return n
	}
	if v < n.value {
		n.left = n.left.Insert(v)
	} else {
		n.right = n.right.Insert(v)
	}

	if n.right.color == red && n.left.color == red {
		if n.right.hasRed() || n.left.hasRed() {
			n.right.color = black
			n.left.color = black
			n.color = red
			return n
		}
	}

	if n.right.color == red && n.right.left.color == red {
		n.right = n.right.rotateRight()
	}
	if n.left.color == red && n.left.right.color == red {
		n.left = n.left.rotateLeft()
	}
	if n.right.color == red && n.right.right.color == red {
		n.color = red
		n.right.color = black
		n.right.right.color = red
		return n.rotateLeft()
	}
	if n.left.color == red && n.left.left.color == red {
		n.color = red
		n.left.color = black
		n.left.left.color = red
		return n.rotateRight()
	}

	return n
}

func (n RBNode) hasRed() bool {
	return n.right.color == red || n.left.color == red
}

func (n RBNode) rotateRight() RBNode {
	parent := n.left
	n.left = parent.right
	parent.right = n
	return parent
}

func (n RBNode) rotateLeft() RBNode {
	parent := n.right
	n.right = parent.left
	parent.left = n
	return parent
}
