package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {

	// Skriv din kode her
	
func Bubble(list []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list) -1; i++{
			if list[i + 1] < list[i]{
				DoSwap(list, i, i + 1)
				swap = true
			}
		}
	}
}
// External swap function for re-use
func DoSwap(list []int, i1 ,i2 int){
	tmp := list[i1];
	list[i1] = list[i2];
	list[i2] = tmp;
}

}

// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QuickSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
