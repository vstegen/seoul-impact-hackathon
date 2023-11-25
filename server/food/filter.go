package food

import (
	"strings"
)

type filter struct {
	// using a pointer here so that we can differentiate between no value and false
	status      RestaurantStatus
	foodOptions map[string]struct{}
}

func (f filter) apply(facilities []Facility) []Facility {
	var filteredFacilities []Facility

	for _, fac := range facilities {
		if f.status != "" && !strings.EqualFold(string(fac.CurrentStatus), string(f.status)) {
			continue
		}

		if f.foodOptions != nil && len(f.foodOptions) > 0 {
			var hasFoodType bool
			for _, foodType := range fac.FoodTypes {
				option := strings.ToLower(string(foodType))
				if _, ok := f.foodOptions[option]; ok {
					hasFoodType = true
					break
				}
			}

			if !hasFoodType {
				continue
			}
		}

		filteredFacilities = append(filteredFacilities, fac)

	}

	return filteredFacilities
}
