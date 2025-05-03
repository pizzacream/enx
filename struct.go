package enx

func GetStruct[T any](k string, d T) T {
	return GetJSON(k, d)
}

func MustGetStruct[T any](k string) T {
	return MustGetJSON[T](k)
}
