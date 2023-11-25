package food

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Repo interface {
	Get() ([]Facility, error)
	GetById(id int) (*Facility, error)
	GetBy(f filter) ([]Facility, error)
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
