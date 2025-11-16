package myfunc

func AddCity(sliceCities []string, city string) []string {
	sliceCities = append(sliceCities, city)
	return sliceCities
}
