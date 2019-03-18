package main

func quick_sort(list []int, start, end int) {

	key := list[start]
	i,j := start, end

	for i <= j {
		for list[j] > key {
			j--
		}


	}
}

func main() {
	source_list := []int{9,4,3,7,11,12,5}

	strart := 0
	end := len(source_list)-1

	quick_sort(source_list, strart, end)

}