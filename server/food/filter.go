package food

import (
	"strings"
)

type filter struct {
	// using a pointer here so that we can differentiate between no value and false
	status RestaurantStatus
}

func (f filter) apply(facilities []Facility) []Facility {
	var filteredFacilities []Facility

	for _, fac := range facilities {
		if f.status != "" && !strings.EqualFold(string(fac.CurrentStatus), string(f.status)) {
			continue
		}

		filteredFacilities = append(filteredFacilities, fac)
	}

	return filteredFacilities
}
