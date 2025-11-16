package myfunc

func DeleteCity(sliceCities []string, delCity string) []string {
	for i, city := range sliceCities {
		if city == delCity {
			sliceCities = append(sliceCities[:i], sliceCities[i+1:]...)
		}
	}
	return sliceCities
}
