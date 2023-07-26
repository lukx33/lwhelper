package lwhelper

func GetKeysFromMap(m map[string]bool) []string {
	list := []string{}
	for k := range m {
		list = append(list, k)
	}
	return list
}
