package food

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"youth-korea/utils"
)

var testFacilities = []Facility{
	{
		Id:            1,
		Name:          "Seoul Seafood",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       utils.String("82-2-3156-5400"),
		Website:       utils.String("enghotline.cafe24.com"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
	},
	{
		Id:            2,
		Name:          "Gukbap Alley",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Contact:       utils.String("+82-2-6364-4000"),
		Website:       utils.String("en.salvationarmy.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusClosed,
	},
	{
		Id:            3,
		Name:          "7-Eleven",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       utils.String("82-2-3156-5400"),
		Website:       utils.String("enghotline.cafe24.com"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
		FoodTypes:     []FoodOption{FoodOptionVegan},
	},
	{
		Id:            4,
		Name:          "Salvation Army Korea",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Contact:       utils.String("+82-2-6364-4000"),
		Website:       utils.String("en.salvationarmy.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
		FoodTypes:     []FoodOption{FoodOptionHalal, FoodOptionVegetarian},
	},
}

func TestGetAllFacilities(t *testing.T) {
	routes := NewRoutes(FacilityRepo{})
	facilityString, _ := json.Marshal(testFacilities)

	setupTestFacilities(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/food", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetFacilities(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(facilityString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetFiltersFacilities(t *testing.T) {
	routes := NewRoutes(FacilityRepo{})

	setupTestFacilities(t)

	expectedFacilities := []Facility{
		{
			Id:            4,
			Name:          "Salvation Army Korea",
			Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
			Contact:       utils.String("+82-2-6364-4000"),
			Website:       utils.String("en.salvationarmy.kr"),
			OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
			CurrentStatus: StatusOpen,
			FoodTypes:     []FoodOption{FoodOptionHalal, FoodOptionVegetarian},
		},
	}
	expectedFacilitiesString, _ := json.Marshal(expectedFacilities)

	e := echo.New()
	q := make(url.Values)
	q.Set("status", "open")
	q.Set("foodOptions", "halal,vegetarian")
	req := httptest.NewRequest(http.MethodGet, "/food/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetFacilities(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedFacilitiesString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestEmptyFiltersFacility(t *testing.T) {
	routes := NewRoutes(FacilityRepo{})

	setupTestFacilities(t)

	expectedFacilities := []Facility{}
	expectedFacilitiesString, _ := json.Marshal(expectedFacilities)

	e := echo.New()
	q := make(url.Values)
	q.Set("status", "invalid")
	req := httptest.NewRequest(http.MethodGet, "/food/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetFacilities(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedFacilitiesString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetShelterByValidId(t *testing.T) {
	routes := NewRoutes(FacilityRepo{})

	setupTestFacilities(t)

	expectedFacilities := testFacilities[1] // ID = 2
	expectedFacilitiesString, _ := json.Marshal(expectedFacilities)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/food", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, routes.GetFacility(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedFacilitiesString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetShelterByInvalidId(t *testing.T) {
	routes := NewRoutes(FacilityRepo{})

	setupTestFacilities(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/food/:id")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := routes.GetFacility(c)
	if assert.Error(t, err) {
		assert.Equal(t, err.(*echo.HTTPError).Code, http.StatusNotFound)
	}
}

func setupTestFacilities(t *testing.T) {
	oldFacilities := facilities
	facilities = testFacilities
	t.Cleanup(func() {
		facilities = oldFacilities
	})
}
