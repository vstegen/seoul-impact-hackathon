package shelters

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

var testShelters = []Shelter{
	{
		Id:            1,
		Name:          "Korea Women’s Hotline",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       utils.String("82-2-3156-5400"),
		Website:       utils.String("enghotline.cafe24.com"),
		Requirements:  utils.String("Under 18 years old"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
	{
		Id:            2,
		Name:          "Salvation Army Korea",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Facilities:    &[]string{"Childcare", "Physical Health", "Education", "Rent Assistance", "Utility Assistance", "Meals", "Emergency Shelters", "Family Shelters", "Youth Shelters", "Counseling", "Clothing"},
		Contact:       utils.String("+82-2-6364-4000"),
		Website:       utils.String("en.salvationarmy.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   15,
		Capacity:      0,
		CurrentStatus: StatusClosed,
	},
	{
		Id:            3,
		Name:          "Anna’s House",
		Address:       "118 Hadaewon-dong, Jungian-gu Seongnam-si, Gyeonggi-do",
		Facilities:    &[]string{"Emergency Shelters", "Meals", "Clothing", "Physical Health", "Counseling", "Education"},
		Contact:       utils.String("031-757-6336"),
		Website:       utils.String("annahouse.or.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   6,
		Capacity:      6,
		CurrentStatus: StatusOpen,
		Announcement:  utils.String("We are currently full, please check back later"),
	},
}

func TestGetAllShelters(t *testing.T) {
	routes := NewRoutes(ShelterRepo{})
	shelterString, _ := json.Marshal(testShelters)

	setupTestShelters(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/shelters", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetShelters(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(shelterString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetFiltersShelters(t *testing.T) {
	routes := NewRoutes(ShelterRepo{})

	setupTestShelters(t)

	expectedShelters := []Shelter{
		{
			Id:            1,
			Name:          "Korea Women’s Hotline",
			Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
			Contact:       utils.String("82-2-3156-5400"),
			Website:       utils.String("enghotline.cafe24.com"),
			Requirements:  utils.String("Under 18 years old"),
			OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
			Rules:         "<Imagine some rules here>",
			MaxCapacity:   20,
			Capacity:      5,
			CurrentStatus: StatusOpen,
		},
	}
	expectedSheltersString, _ := json.Marshal(expectedShelters)

	e := echo.New()
	q := make(url.Values)
	q.Set("status", "open")
	q.Set("hasCapacity", "true")
	req := httptest.NewRequest(http.MethodGet, "/shelters/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetShelters(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedSheltersString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestEmptyFiltersShelters(t *testing.T) {
	routes := NewRoutes(ShelterRepo{})

	setupTestShelters(t)

	expectedShelters := []Shelter{}
	expectedSheltersString, _ := json.Marshal(expectedShelters)

	e := echo.New()
	q := make(url.Values)
	q.Set("status", "full")
	req := httptest.NewRequest(http.MethodGet, "/shelters/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, routes.GetShelters(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedSheltersString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetShelterByValidId(t *testing.T) {
	routes := NewRoutes(ShelterRepo{})

	setupTestShelters(t)

	expectedShelters := testShelters[1] // ID = 2
	expectedSheltersString, _ := json.Marshal(expectedShelters)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/shelters", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, routes.GetShelter(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedSheltersString), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetShelterByInvalidId(t *testing.T) {
	routes := NewRoutes(ShelterRepo{})

	setupTestShelters(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/shelters/:id")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := routes.GetShelter(c)
	if assert.Error(t, err) {
		assert.Equal(t, err.(*echo.HTTPError).Code, http.StatusNotFound)
	}
}

func setupTestShelters(t *testing.T) {
	oldShelters := shelters
	shelters = testShelters
	t.Cleanup(func() {
		shelters = oldShelters
	})
}
