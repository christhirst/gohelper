package iterables

func FindAndDelete(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func GetMapKeys[K any](m map[string]K) []string {
	i := 0
	var keys []string
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
