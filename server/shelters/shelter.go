package shelters

type (
	ShelterStatus string
)

const (
	StatusOpen   ShelterStatus = "Open"
	StatusClosed ShelterStatus = "Closed"
	StatusFull   ShelterStatus = "Full"
)

type Shelter struct {
	Id            int           `json: "id"`
	OpeningTime   string        `json: "openingTime"`
	Facilities    []string      `json: "facilities", omitempty`
	Requirements  string        `json: "requirements", omitempty`
	Rules         string        `json: "rules"`
	MaxCapacity   int           `json: "maxCapacity"`
	Capacity      int           `json: "capacity"`
	Rating        int           `json: "rating", omitempty`
	Desc          string        `json: "desc", omitempty`
	Address       string        `json: "address"`
	Name          string        `json: "name"`
	CurrentStatus ShelterStatus `json: "currentStatus"`
	Announcement  string        `json: "announcement", omitempty`
	Contact       string        `json: "contact", omitempty`
	Website       string        `json: "website", omitempty`
}

func (s *Shelter) hasCapacity() bool {
	return s.Capacity < s.MaxCapacity && s.CurrentStatus != StatusFull
}
