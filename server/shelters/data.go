package shelters

import (
	"errors"
	"fmt"

	"youth-korea/utils"
)

var errNotFound = errors.New("shelter not found")

var shelters = []Shelter{
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
		CurrentStatus: StatusFull,
		Announcement:  utils.String("We are currently full, please check back later"),
	},
	{
		Id:            4,
		Name:          "SOS Children's Village Seoul",
		Address:       "Seoul, South Korea",
		Facilities:    &[]string{"Childcare", "Youth Shelters", "Meals", "Physical Health", "Education"},
		Contact:       utils.String("02-3453-8400"),
		Website:       utils.String("sos-childrensvillages.org"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
	{
		Id:            5,
		Name:          "Happy Home",
		Address:       "675-2 Bupyeong 2(i)-dong, Bupyeong-gu, Incheon, South Korea",
		Facilities:    &[]string{"Childcare", "Youth Shelters", "Meals", "Education", "Counseling"},
		Contact:       utils.String("+82 32-518-2080"),
		Website:       utils.String("happyhome.or.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      20,
		CurrentStatus: StatusFull,
	},
	{
		Id:            6,
		Name:          "Rainbow Youth Center",
		Address:       "20 Jahamun-ro 24-gil, Jongno-gu, Seoul, South Korea",
		Facilities:    &[]string{"Youth Shelters", "Counseling", "Education"},
		Contact:       utils.String("02-733-7587"),
		Website:       utils.String("rainbowyouth.or.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
}

type ShelterRepo struct{}

// NOTE: It is intended for this method to always return a nil error.
// This is done for extensibility later when the underlying data source
// changes into a potentially failing one.
func (r ShelterRepo) Get() ([]Shelter, error) {
	return shelters, nil
}

func (r ShelterRepo) GetById(id int) (*Shelter, error) {
	for _, s := range shelters {
		if s.Id == id {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("Shelter with id %d not found", id)
}

func (r ShelterRepo) GetBy(f filter) ([]Shelter, error) {
	shelters, err := r.Get()
	if err != nil {
		return nil, errNotFound
	}

	return f.apply(shelters), nil
}
