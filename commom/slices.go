package commom

func PartationSlice[T any](s []T, step int) [][]T {
	res := make([][]T, 0)
	if len(s) < step {
		res = append(res, s)
		return res
	}
	for i := 0; i < len(s); i += step {
		end := i + step
		if end < len(s) {
			res = append(res, s[i:end])
		} else {
			res = append(res, s[i:])
		}
	}
	return res
}
