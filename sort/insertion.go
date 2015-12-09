package sort

func Insertion(data Interface) {
	len := data.Len()

	for i := 0; i < len; i++ {
	inner:

		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break inner
			}
		}
	}
}
