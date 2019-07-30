package arrutil

/*
IntIndexOf finds the index of target element
it will return -1 while find nothing
*/
func IntIndexOf(elements []int, targetElement int) (index int) {
	for index, element := range elements {
		if element == targetElement {
			return index
		}
	}
	return -1
}
