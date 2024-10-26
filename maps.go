package lwhelper

import "sort"

func GetMapFromStrinSlice(ss []string) map[string]bool {
	res := map[string]bool{}
	for _, k := range ss {
		res[k] = true
	}
	return res
}

func GetKeysFromMap[V any](m map[string]V) []string {
	list := []string{}
	for k := range m {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func GetKeysFromMapInt[V any](m map[int]V) []int {
	list := []int{}
	for k := range m {
		list = append(list, k)
	}
	sort.Ints(list)
	return list
}
