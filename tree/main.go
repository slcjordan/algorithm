package main

import "fmt"
import "strings"

func printout(root RBNode) {
	fmt.Println()
	fmt.Println(root)
	fmt.Println()
}

func main() {
	root := RBNode{}
	for _, word := range strings.Split(
		`the quick brown fox jumped over the lazy dog on his way to lazy town where he was said to bed the night with a fair maiden and a fierce dragon grrr`,
		" ") {
		root = root.Insert(word)
	}
	printout(root)
}
