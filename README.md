# seoul-impact-hackathon

## Server

Build the server:

```
cd ./server
go build
```

Run the server:

```
./youth-korea
```

### API endpoints

The server is started on port `:8000`

- `/shelters`: Returns all shelters irrespective of status
  A shelter adheres to the following format:

```go

type Shelter struct {
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

```

