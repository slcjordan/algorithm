package sort

var maxInt = int(^uint(0) >> 1)

// merge will take two channels in sorted order and return a channel with sorted order.
func merge(a <-chan int, b <-chan int, out chan<- int) {
	var lastA, lastB int
	var aOpen, bOpen bool
	aGate := a
	bGate := b

	for {
		select {
		case lastA, aOpen = <-aGate:
			aGate = nil
			if !aOpen {
				drain(b, out, lastB, bOpen)
				return
			}
		case lastB, bOpen = <-bGate:
			bGate = nil
			if !bOpen {
				drain(a, out, lastA, aOpen)
				return
			}
		}
		if !(aOpen && bOpen) {
			continue
		}

		if lastA < lastB {
			out <- lastA
			aGate = a
		} else {
			out <- lastB
			bGate = b
		}
	}
}

func drain(in <-chan int, out chan<- int, first int, cond bool) {
	if cond {
		out <- first
	}

	for n := range in {
		out <- n
	}
}

// MergeSort will take an array and return a channel of those values in sorted order.
func MergeSort(x []int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		if len(x) <= 1 {
			for _, n := range x {
				result <- n
			}
			return
		}

		mid := len(x) / 2
		a := MergeSort(x[:mid])
		b := MergeSort(x[mid:])
		merge(a, b, result)
	}()
	return result
}
