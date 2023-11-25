package shelters

import (
	"strings"
)

type filter struct {
	// using a pointer here so that we can differentiate between no value and false
	hasCapacity *bool
	status      ShelterStatus
}

func (f filter) apply(shelters []Shelter) []Shelter {
	var filteredShelters []Shelter

	for _, shelter := range shelters {
		if f.status != "" && !strings.EqualFold(string(shelter.CurrentStatus), string(f.status)) {
			continue
		}

		if f.hasCapacity != nil && *f.hasCapacity && !shelter.hasCapacity() {
			continue
		}

		filteredShelters = append(filteredShelters, shelter)
	}

	return filteredShelters
}
