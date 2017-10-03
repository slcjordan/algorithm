package main

import "fmt"
import "strings"

func printout(n RBNode, depth int) {
	if n.initialized != nil {
		printout(n.left, depth+1)
		msg := fmt.Sprintf("%% %ds%%s %%s\n", depth*4)
		if n.color {
			fmt.Printf(msg, "", n.value, "(red)")
		} else {
			fmt.Printf(msg, "", n.value, "(black)")
		}
		printout(n.right, depth+1)
	}
}

func main() {
	root := RBNode{}
	for _, word := range strings.Split(
		`the quick brown fox jumped over the lazy dog on his way to lazy town where he was said to woo fair maiden and battle a fierce dragon grrr. Maybe one day he will be king`,
		" ") {
		root = root.Insert(word)
	}
	printout(root, 0)
}
