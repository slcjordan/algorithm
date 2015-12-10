package sort

func Merge(data []string) <-chan string {
	if len(data) > 1 {
		mid := len(data) / 2
		firstHalf := data[:mid]
		secondHalf := data[mid:]
		return combine(Merge(firstHalf), Merge(secondHalf), firstHalf, secondHalf)
	}

	result := make(chan string)

	go func() {
		defer close(result)

		for _, x := range data {
			result <- x
		}
	}()
	return result
}

func combine(a, b <-chan string, firstHalf, secondHalf []string) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)

		pickX := true
		pickY := true
		var x, y string
		var ok, onePass bool

		for {
			if pickX {
				x, ok = <-a
				if !ok {
					if onePass {
						result <- y
						for y = range b {
							result <- y
						}
					}
					return
				}
			}
			if pickY {
				y, ok = <-b
				if !ok {
					if onePass {
						result <- x
						for x = range a {
							result <- x
						}
					}
					return
				}
			}
			pickX = false
			pickY = false

			if x < y {
				result <- x
				pickX = true
			} else {
				result <- y
				pickY = true
			}
			onePass = true
		}
	}()

	return result
}
