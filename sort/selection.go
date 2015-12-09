package sort

func Selection(data Interface) {
	len := data.Len()

	for i := 0; i < len; i++ {
		min := i

		for j := i; j < len; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
