package shelters

import (
	"strings"
)

type baseFilter struct {
	// using a pointer here so that we can differentiate between no value and the default false
	hasCapacity *bool
	status      ShelterStatus
}

type filter interface {
	apply(shelters []Shelter) []Shelter
}

func (f baseFilter) apply(shelters []Shelter) []Shelter {
	var filteredShelters []Shelter

	for _, shelter := range shelters {
		if f.status != "" && !strings.EqualFold(string(shelter.CurrentStatus), string(f.status)) {
			continue
		}

		if f.hasCapacity != nil && *f.hasCapacity && !shelter.hasCapacity() {
			continue
		}

		if f.hasCapacity != nil && !*f.hasCapacity && shelter.hasCapacity() {
			continue
		}

		filteredShelters = append(filteredShelters, shelter)
	}

	return filteredShelters
}
