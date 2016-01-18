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

func TestMerge(t *testing.T) {
	result := MergeSort([]int{
		9,
		8,
		7,
		6,
		4,
		5,
		3,
		1,
		2,
		0,
	})

	count := 0
	for n := range result {
		if n != count {
			t.Fatalf("expected %d got %d", count, n)
		}
		count++
	}
	if count != 10 {
		t.Fatalf("expected 10 numbers to be passed in result channel got %d", count)
	}
}
