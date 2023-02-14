package binarysearch

import "sort"

func SearchInts(list []int, key int) int {
	i := sort.Search(len(list), func(i int) bool { return list[i] >= key })
	if i < len(list) && list[i] == key {
		return i
	} else {
		return -1
	}
}
