package sort

import (
	std "sort"
	"testing"
)

func TestAll(t *testing.T) {
	for name, algorithm := range map[string]func(Interface){
		"Selection": Selection,
		"Insertion": Insertion,
	} {
		data := std.StringSlice([]string{
			"zero",
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
			"ten",
		})
		sorted := []string{
			"eight",
			"five",
			"four",
			"nine",
			"one",
			"seven",
			"six",
			"ten",
			"three",
			"two",
			"zero",
		}
		algorithm(data)

		for i := range sorted {
			if data[i] != sorted[i] {
				t.Fatal(name, data)
			}
		}
	}
}
