package must

func Must[T any, E comparable](result T, err E) T {
	var e E
	if err != e {
		panic(err)
	}
	return result
}
