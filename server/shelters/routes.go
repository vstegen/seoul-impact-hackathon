package shelters

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TODO: potentially add query param validation for the filter
func GetShelters(c echo.Context) error {
	activeFilter := filter{}
	activeFilter.status = ShelterStatus(c.QueryParam("status"))

	hasCapacityQueryParam := c.QueryParam("hasCapacity")
	if hasCapacityQueryParam != "" {
		hasCapacity, err := strconv.ParseBool(hasCapacityQueryParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid hasCapacity '%s': expected true or false", hasCapacityQueryParam))
		}
		activeFilter.hasCapacity = &hasCapacity
	}

	filteredShelters := activeFilter.apply(shelters)
	if len(filteredShelters) == 0 {
		return c.JSON(http.StatusOK, []Shelter{})
	}

	return c.JSON(http.StatusOK, filteredShelters)
}

func GetShelter(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid id '%s': %w", idParam, err))
	}
	// TODO: use a map instead of looping if significantly more shelters are added
	for _, shelter := range shelters {
		if id == shelter.Id {
			return c.JSON(http.StatusOK, shelter)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, nil)
}
