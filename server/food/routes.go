package food

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type repo interface {
	Get() ([]Facility, error)
	GetById(id int) (*Facility, error)
	GetBy(f filter) ([]Facility, error)
}

// Routes is a collection of the handlers available for facilities.
type Routes struct {
	repo repo
}

// NewRoutes returns a new instance of Routes that configures
// the given repo as the underlying data source.
func NewRoutes(r repo) Routes {
	return Routes{
		repo: r,
	}
}

// GetFacilities returns a list of facilities.
// If no query parameters are given, it returns all facilities,
// otherwise it will filter the facilities based on the given parameters.
//
// Accepted query parameters:
// - status: string
// - foodOptions: comma-separated list of strings, e.g. "vegan,vegetarian"
func (r Routes) GetFacilities(c echo.Context) error {
	activeFilter := filter{}
	activeFilter.status = RestaurantStatus(c.QueryParam("status"))

	foodOptionsQueryParam := c.QueryParam("foodOptions")
	if foodOptionsQueryParam != "" {
		foodOptions := strings.Split(foodOptionsQueryParam, ",")
		activeFilter.foodOptions = make(map[string]struct{})
		for _, foodOption := range foodOptions {
			activeFilter.foodOptions[strings.ToLower(foodOption)] = struct{}{}
		}
	}

	filteredFacilities, err := r.repo.GetBy(activeFilter)
	if err != nil {
		if errors.Is(err, errNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error getting facility: %w", err))
		}
	}

	if len(filteredFacilities) == 0 {
		return c.JSON(http.StatusOK, []Facility{})
	}

	return c.JSON(http.StatusOK, filteredFacilities)
}

// GetShelter returns a single shelter.
// It makes the assumption that the URI contains the ID of the shelter.
func (r Routes) GetFacility(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid id '%s': %w", idParam, err))
	}

	fac, err := r.repo.GetById(id)
	if err != nil {
		if errors.Is(err, errNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error getting facility: %w", err))
		}
	}

	return c.JSON(http.StatusOK, fac)
}
