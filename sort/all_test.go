package sort

import (
	std "sort"
	"testing"
)

func TestInPlace(t *testing.T) {
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

func TestChan(t *testing.T) {
	for name, algorithm := range map[string]func([]string) <-chan string{
		"Merge": Merge,
	} {
		data := []string{
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
		}
		expected := []string{
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

		i := 0
		sorted := make([]string, 0, len(expected))

		for x := range algorithm(data) {
			sorted = append(sorted, x)

			if expected[i] != x {
				t.Fatal(i, name, sorted, "\""+x+"\"")
			}
			i += 1
		}
		if len(sorted) != len(expected) {
			t.Fatal("bad length", name, sorted)
		}
	}
	_ = "breakpoint"
}
