package solution

/*
problem:
Given an array of ints = [6, 4, -3, 5, -2, -1, 0, 1, -9],
implement a function in any language to move all positive numbers to the left,
and move all non-positive numbers to the right.
Try your best to make its time complexity to O(n), and space complexity to O(1).
*/

func splitArray(a []int) {
	if a == nil {
		return
	}

	const splitter = 0

	head, tail := 0, len(a)-1

	for i := 1; i <= tail; {
		if a[i] < splitter {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[head] = a[head], a[i]
			head++
			i++
		}
	}

}
