package shelters

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Repo interface {
	Get() ([]Shelter, error)
	GetById(id int) (*Shelter, error)
	GetBy(f filter) ([]Shelter, error)
}

type Routes struct {
	repo Repo
}

func NewRoutes(r Repo) Routes {
	return Routes{
		repo: r,
	}
}

// TODO: potentially add query param validation for the filter
func (r Routes) GetShelters(c echo.Context) error {
	activeFilter := baseFilter{}
	activeFilter.status = ShelterStatus(c.QueryParam("status"))

	hasCapacityQueryParam := c.QueryParam("hasCapacity")
	if hasCapacityQueryParam != "" {
		hasCapacity, err := strconv.ParseBool(hasCapacityQueryParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid hasCapacity '%s': expected true or false", hasCapacityQueryParam))
		}
		activeFilter.hasCapacity = &hasCapacity
	}

	filteredShelters, err := r.repo.GetBy(activeFilter)
	if err != nil {
		if errors.Is(err, errNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error getting shelter: %w", err))
		}
	}

	if len(filteredShelters) == 0 {
		return c.JSON(http.StatusOK, []Shelter{})
	}

	return c.JSON(http.StatusOK, filteredShelters)
}

func (r Routes) GetShelter(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid id '%s': %w", idParam, err))
	}

	shelter, err := r.repo.GetById(id)
	if err != nil {
		if errors.Is(err, errNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error getting shelter: %w", err))
		}
	}

	return c.JSON(http.StatusOK, shelter)
}
