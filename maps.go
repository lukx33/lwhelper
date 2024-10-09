package lwhelper

func GetKeysFromMap[V any](m map[string]V) []string {
	list := []string{}
	for k := range m {
		list = append(list, k)
	}
	return list
}
