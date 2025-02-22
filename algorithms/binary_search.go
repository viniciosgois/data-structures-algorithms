package algorithms

func BinarySearch(list []int, target int) (int, int) {
	lower := 0
	higher := len(list) - 1

	for lower <= higher {
		middle := (lower + higher) / 2
		prob := list[middle]

		if prob == target {
			return prob, middle
		}

		if prob < target {
			lower = middle + 1
		} else {
			higher = middle - 1
		}
	}

	return 0, 0
}
